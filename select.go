package cql

import (
	"cql/utils"
	"strings"
	"sync"
)

type Select struct {
	column string
	from   string
	join   string
	where  string
	group  string
	having string
	order  string
	offset int64
	limit  int64
	Values []interface{}
	rwm    sync.RWMutex
}

func NewSelect() *Select {
	return &Select{
		column: "*",
	}
}
func (s *Select) Column(c ...string) *Select {
	s.rwm.Lock()
	if s.column == "*" && len(c) > 0 {
		s.column = strings.Join(c, ", ")
	} else if s.column != "*" && len(c) > 0 {
		s.column += ", " + strings.Join(c, ", ")
	}
	s.rwm.Unlock()
	return s
}

func (s *Select) From(f string) *Select {
	s.rwm.Lock()
	s.from = f
	s.rwm.Unlock()
	return s
}

func (s *Select) Join(j ...string) *Select {
	s.rwm.Lock()
	s.join = strings.Join(j, "JOIN ")
	s.rwm.Unlock()
	return s
}

func (s *Select) Where(w map[string]map[string]map[string]interface{}) *Select {
	kw, kv := utils.GetWherev(w)
	s.rwm.Lock()
	s.where = kw
	s.Values = append(s.Values, kv...)
	s.rwm.Unlock()
	return s
}

func (s *Select) Group(g ...string) *Select {
	s.rwm.Lock()
	s.group = strings.Join(g, ", ")
	s.rwm.Unlock()
	return s
}

func (s *Select) Having(h string) *Select {
	s.rwm.Lock()
	s.having = h
	s.rwm.Unlock()
	return s
}

func (s *Select) Order(o ...string) *Select {
	s.rwm.Lock()
	s.order = strings.Join(o, ", ")
	s.rwm.Unlock()
	return s
}

func (s *Select) Offset(o int64) *Select {
	s.rwm.Lock()
	s.offset = o
	s.rwm.Unlock()
	return s
}

func (s *Select) Limit(l int64) *Select {
	s.rwm.Lock()
	s.limit = l
	s.rwm.Unlock()
	return s
}

func (s *Select) String() (sql string) {
	if len(s.from) == 0 {
		return
	}
	s.rwm.Lock()
	desc := []string{
		"SELECT",
		s.column,
		"FROM",
		s.from,
	}

	if len(s.where) > 0 {
		desc = append(desc, "WHERE", s.where)
	}

	if len(s.group) > 0 {
		desc = append(desc, "GROUP BY", s.group)
	}

	if len(s.having) > 0 {
		desc = append(desc, "HAVING", s.having)
	}

	if len(s.order) > 0 {
		desc = append(desc, "ORDER BY", s.order)
	}

	if s.limit > 0 && s.offset > 0 {
		desc = append(desc, "LIMIT", "?,?")
		s.Values = append(s.Values, s.offset, s.limit)
	} else if s.limit > 0 {
		desc = append(desc, "LIMIT", "?")
		s.Values = append(s.Values, s.limit)
	}

	sql = strings.Join(desc, " ")
	s.rwm.Unlock()
	return
}
