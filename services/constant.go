package services

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/dronm/ds/pgds"
	"github.com/dronm/gobizapp/database"
	"github.com/dronm/gobizapp/errs"

	"github.com/dronm/gobizapp/models"
)

var ConstantServiceInstance *ConstantService

func ConstantServiceManager() *ConstantService {
	if ConstantServiceInstance == nil {
		ConstantServiceInstance = NewConstantService(database.DB)
	}
	return ConstantServiceInstance
}

type ValidateConstantValue func(val string) error

type Constant struct {
	SetValue ValidateConstantValue
	RolesSet map[string]struct{}
	RolesGet map[string]struct{}
}

type ConstantService struct {
	DB             *pgds.PgProvider
	mu             sync.RWMutex
	initialized    bool
	constantLookup map[string]*Constant
}

func NewConstantService(db *pgds.PgProvider) *ConstantService {
	return &ConstantService{
		DB:             db,
		constantLookup: make(map[string]*Constant),
	}
}

// Initialize populates the constantLookup map once.
func (s *ConstantService) Initialize(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Prevent double loading
	if s.initialized {
		return nil
	}

	poolConn, connID, err := s.DB.GetSecondary("")
	if err != nil {
		return fmt.Errorf("GetSecondary() failed: %v", err)
	}
	defer s.DB.Release(poolConn, connID)
	conn := poolConn.Conn()

	rows, err := conn.Query(ctx, `SELECT id, roles_set, roles_get FROM constants_list_view`)
	if err != nil {
		return fmt.Errorf("conn.Query() failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var rolesSet []string
		var rolesGet []string
		if err := rows.Scan(&id, &rolesSet, &rolesGet); err != nil {
			return fmt.Errorf("rows.Scan() failed: %w", err)
		}
		constObj := Constant{
			RolesGet: make(map[string]struct{}, len(rolesGet)),
			RolesSet: make(map[string]struct{}, len(rolesSet)),
		}
		for _, roleID := range rolesGet {
			constObj.RolesGet[roleID] = struct{}{}
		}
		for _, roleID := range rolesSet {
			constObj.RolesSet[roleID] = struct{}{}
		}
		s.constantLookup[id] = &constObj
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("rows.Err() failed: %w", err)
	}

	s.initialized = true
	return nil
}

// init is a helper function. It checks and if not initialized,
// starts initialization.
func (s *ConstantService) init(ctx context.Context) error {
	s.mu.RLock()
	initialized := s.initialized
	s.mu.RUnlock()

	if !initialized {
		if err := s.Initialize(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (s *ConstantService) SetValidationFunc(ctx context.Context, id string, f ValidateConstantValue) error {
	constObj, err := s.Get(ctx, id)
	if err != nil {
		return err
	}
	if constObj == nil {
		return errs.NewPublicError("CONST_UNDEFINED")
	}
	constObj.SetValue = f
	return nil
}

func (s *ConstantService) Get(ctx context.Context, id string) (*Constant, error) {
	if err := s.init(ctx); err != nil {
		return nil, err
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	constObj, exists := s.constantLookup[id]
	if exists {
		return constObj, nil
	}
	return nil, errs.NewPublicError("CONST_UNDEFINED")
}

// Exists checks if a constant ID exists, calling Initialize if needed.
func (s *ConstantService) Exists(ctx context.Context, id string) (bool, error) {
	if err := s.init(ctx); err != nil {
		return false, err
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	_, exists := s.constantLookup[id]
	return exists, nil
}

func (s *ConstantService) Update(ctx context.Context, id, val string) error {
	constObj, err := s.Get(ctx, id)
	if err != nil {
		return err
	}
	if constObj == nil {
		return errs.NewPublicError("CONST_UNDEFINED")
	}
	if constObj.SetValue != nil {
		if err := constObj.SetValue(val); err != nil {
			return err
		}
	}

	poolConn, connID, err := s.DB.GetPrimary()
	if err != nil {
		return fmt.Errorf("GetPrimary() failed: %v", err)
	}
	defer s.DB.Release(poolConn, connID)
	conn := poolConn.Conn()

	_, err = conn.Exec(ctx, fmt.Sprintf(`UPDATE const_%s SET val = $1`, id), val)
	if err != nil {
		return fmt.Errorf("conn.Exec() UPDATE failed: %w", err)
	}

	return nil
}

func (s *ConstantService) FetchValues(ctx context.Context, idList []string, userRoleID string) ([]models.ConstantValList, error) {
	s.mu.RLock()
	initialized := s.initialized
	s.mu.RUnlock()

	if !initialized {
		if err := s.Initialize(ctx); err != nil {
			return nil, fmt.Errorf("s.Initialize(): %w", err)
		}
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	var query strings.Builder
	for _, id := range idList {
		_, ok := s.constantLookup[id]
		if !ok {
			return nil, errs.NewPublicErrorWithTemplate("CONST_UNDEFINED", id)
		}
		// permission check
		// if _, ok := constObj.RolesGet[userRoleID]; !ok {
		// 	return nil, errs.NewPublicErrorCustom("NOT_ALLOWED", fmt.Sprintf("get is not allowed for constant %s, role %s", id, userRoleID))
		// }

		if query.Len() > 0 {
			query.WriteString(" UNION ALL ")
		}

		query.WriteString(fmt.Sprintf("SELECT '%s', (SELECT val FROM const_%s)", id, id))
	}

	poolConn, connID, err := s.DB.GetSecondary("")
	if err != nil {
		return nil, fmt.Errorf("GetSecondary() failed: %v", err)
	}
	defer s.DB.Release(poolConn, connID)
	conn := poolConn.Conn()

	var modelList []models.ConstantValList

	rows, err := conn.Query(ctx, query.String())
	if err != nil {
		return nil, fmt.Errorf("conn.Query() SELECT constants: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		m := models.ConstantValList{}

		err := rows.Scan(&m.ID, &m.Val)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan(): %w", err)
		}
		modelList = append(modelList, m)
	}

	return modelList, nil
}
