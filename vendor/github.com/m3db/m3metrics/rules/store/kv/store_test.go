// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package kv

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/m3db/m3cluster/kv/mem"
	"github.com/m3db/m3metrics/generated/proto/schema"
	"github.com/m3db/m3metrics/rules"

	"github.com/stretchr/testify/require"
)

const (
	testNamespaceKey  = "testKey"
	testNamespace     = "fooNs"
	testRuleSetKeyFmt = "rules/%s"
)

var (
	testNamespaces = &schema.Namespaces{
		Namespaces: []*schema.Namespace{
			&schema.Namespace{
				Name: "fooNs",
				Snapshots: []*schema.NamespaceSnapshot{
					&schema.NamespaceSnapshot{
						ForRulesetVersion: 1,
						Tombstoned:        false,
					},
					&schema.NamespaceSnapshot{
						ForRulesetVersion: 2,
						Tombstoned:        false,
					},
				},
			},
			&schema.Namespace{
				Name: "barNs",
				Snapshots: []*schema.NamespaceSnapshot{
					&schema.NamespaceSnapshot{
						ForRulesetVersion: 1,
						Tombstoned:        false,
					},
					&schema.NamespaceSnapshot{
						ForRulesetVersion: 2,
						Tombstoned:        true,
					},
				},
			},
		},
	}

	testRuleSetKey = fmt.Sprintf(testRuleSetKeyFmt, testNamespace)
	testRuleSet    = &schema.RuleSet{
		Uuid:               "ruleset",
		Namespace:          "fooNs",
		CreatedAtNanos:     1234,
		LastUpdatedAtNanos: 5678,
		Tombstoned:         false,
		CutoverNanos:       34923,
		MappingRules: []*schema.MappingRule{
			&schema.MappingRule{
				Uuid: "12669817-13ae-40e6-ba2f-33087b262c68",
				Snapshots: []*schema.MappingRuleSnapshot{
					&schema.MappingRuleSnapshot{
						Name:         "foo",
						Tombstoned:   false,
						CutoverNanos: 12345,
						Filter:       "tag1:value1 tag2:value2",
						Policies: []*schema.Policy{
							&schema.Policy{
								StoragePolicy: &schema.StoragePolicy{
									Resolution: &schema.Resolution{
										WindowSize: int64(10 * time.Second),
										Precision:  int64(time.Second),
									},
									Retention: &schema.Retention{
										Period: int64(24 * time.Hour),
									},
								},
								AggregationTypes: []schema.AggregationType{
									schema.AggregationType_P999,
								},
							},
						},
					},
					&schema.MappingRuleSnapshot{
						Name:         "foo",
						Tombstoned:   false,
						CutoverNanos: 67890,
						Filter:       "tag3:value3 tag4:value4",
						Policies: []*schema.Policy{
							&schema.Policy{
								StoragePolicy: &schema.StoragePolicy{
									Resolution: &schema.Resolution{
										WindowSize: int64(time.Minute),
										Precision:  int64(time.Minute),
									},
									Retention: &schema.Retention{
										Period: int64(24 * time.Hour),
									},
								},
							},
							&schema.Policy{
								StoragePolicy: &schema.StoragePolicy{
									Resolution: &schema.Resolution{
										WindowSize: int64(5 * time.Minute),
										Precision:  int64(time.Minute),
									},
									Retention: &schema.Retention{
										Period: int64(48 * time.Hour),
									},
								},
							},
						},
					},
				},
			},
			&schema.MappingRule{
				Uuid: "12669817-13ae-40e6-ba2f-33087b262c68",
				Snapshots: []*schema.MappingRuleSnapshot{
					&schema.MappingRuleSnapshot{
						Name:         "dup",
						Tombstoned:   false,
						CutoverNanos: 12345,
						Filter:       "tag1:value1 tag2:value2",
						Policies: []*schema.Policy{
							&schema.Policy{
								StoragePolicy: &schema.StoragePolicy{
									Resolution: &schema.Resolution{
										WindowSize: int64(10 * time.Second),
										Precision:  int64(time.Second),
									},
									Retention: &schema.Retention{
										Period: int64(24 * time.Hour),
									},
								},
								AggregationTypes: []schema.AggregationType{
									schema.AggregationType_P999,
								},
							},
						},
					},
				},
			},
		},
		RollupRules: []*schema.RollupRule{
			&schema.RollupRule{
				Uuid: "12669817-13ae-40e6-ba2f-33087b262c68",
				Snapshots: []*schema.RollupRuleSnapshot{
					&schema.RollupRuleSnapshot{
						Name:         "foo2",
						Tombstoned:   false,
						CutoverNanos: 12345,
						Filter:       "tag1:value1 tag2:value2",
						Targets: []*schema.RollupTarget{
							&schema.RollupTarget{
								Name: "rName1",
								Tags: []string{"rtagName1", "rtagName2"},
								Policies: []*schema.Policy{
									&schema.Policy{
										StoragePolicy: &schema.StoragePolicy{
											Resolution: &schema.Resolution{
												WindowSize: int64(10 * time.Second),
												Precision:  int64(time.Second),
											},
											Retention: &schema.Retention{
												Period: int64(24 * time.Hour),
											},
										},
									},
								},
							},
						},
					},
					&schema.RollupRuleSnapshot{
						Name:         "bar",
						Tombstoned:   true,
						CutoverNanos: 67890,
						Filter:       "tag3:value3 tag4:value4",
						Targets: []*schema.RollupTarget{
							&schema.RollupTarget{
								Name: "rName1",
								Tags: []string{"rtagName1", "rtagName2"},
								Policies: []*schema.Policy{
									&schema.Policy{
										StoragePolicy: &schema.StoragePolicy{
											Resolution: &schema.Resolution{
												WindowSize: int64(time.Minute),
												Precision:  int64(time.Minute),
											},
											Retention: &schema.Retention{
												Period: int64(24 * time.Hour),
											},
										},
									},
									&schema.Policy{
										StoragePolicy: &schema.StoragePolicy{
											Resolution: &schema.Resolution{
												WindowSize: int64(5 * time.Minute),
												Precision:  int64(time.Minute),
											},
											Retention: &schema.Retention{
												Period: int64(48 * time.Hour),
											},
										},
										AggregationTypes: []schema.AggregationType{
											schema.AggregationType_MEAN,
										},
									},
								},
							},
						},
					},
				},
			},
			&schema.RollupRule{
				Uuid: "12669817-13ae-40e6-ba2f-33087b262c68",
				Snapshots: []*schema.RollupRuleSnapshot{
					&schema.RollupRuleSnapshot{
						Name:         "foo",
						Tombstoned:   false,
						CutoverNanos: 12345,
						Filter:       "tag1:value1 tag2:value2",
						Targets: []*schema.RollupTarget{
							&schema.RollupTarget{
								Name: "rName1",
								Tags: []string{"rtagName1", "rtagName2"},
								Policies: []*schema.Policy{
									&schema.Policy{
										StoragePolicy: &schema.StoragePolicy{
											Resolution: &schema.Resolution{
												WindowSize: int64(10 * time.Second),
												Precision:  int64(time.Second),
											},
											Retention: &schema.Retention{
												Period: int64(24 * time.Hour),
											},
										},
									},
								},
							},
						},
					},
					&schema.RollupRuleSnapshot{
						Name:         "baz",
						Tombstoned:   false,
						CutoverNanos: 67890,
						Filter:       "tag3:value3 tag4:value4",
						Targets: []*schema.RollupTarget{
							&schema.RollupTarget{
								Name: "rName1",
								Tags: []string{"rtagName1", "rtagName2"},
								Policies: []*schema.Policy{
									&schema.Policy{
										StoragePolicy: &schema.StoragePolicy{
											Resolution: &schema.Resolution{
												WindowSize: int64(time.Minute),
												Precision:  int64(time.Minute),
											},
											Retention: &schema.Retention{
												Period: int64(24 * time.Hour),
											},
										},
									},
									&schema.Policy{
										StoragePolicy: &schema.StoragePolicy{
											Resolution: &schema.Resolution{
												WindowSize: int64(5 * time.Minute),
												Precision:  int64(time.Minute),
											},
											Retention: &schema.Retention{
												Period: int64(48 * time.Hour),
											},
										},
										AggregationTypes: []schema.AggregationType{
											schema.AggregationType_MEAN,
										},
									},
								},
							},
						},
					},
				},
			},
			&schema.RollupRule{
				Uuid: "12669817-13ae-40e6-ba2f-33087b262c68",
				Snapshots: []*schema.RollupRuleSnapshot{
					&schema.RollupRuleSnapshot{
						Name:         "dup",
						Tombstoned:   false,
						CutoverNanos: 12345,
						Filter:       "tag1:value1 tag2:value2",
						Targets: []*schema.RollupTarget{
							&schema.RollupTarget{
								Name: "rName1",
								Tags: []string{"rtagName1", "rtagName2"},
								Policies: []*schema.Policy{
									&schema.Policy{
										StoragePolicy: &schema.StoragePolicy{
											Resolution: &schema.Resolution{
												WindowSize: int64(10 * time.Second),
												Precision:  int64(time.Second),
											},
											Retention: &schema.Retention{
												Period: int64(24 * time.Hour),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
)

func TestRuleSetKey(t *testing.T) {
	s := testStore()
	defer s.Close()

	key := s.(*store).ruleSetKey(testNamespace)
	require.Equal(t, "rules/fooNs", key)
}

func TestNewStore(t *testing.T) {
	opts := NewStoreOptions(testNamespaceKey, testRuleSetKeyFmt, nil)
	kvStore := mem.NewStore()
	s := NewStore(kvStore, opts).(*store)
	defer s.Close()

	require.Equal(t, s.kvStore, kvStore)
	require.Equal(t, s.opts, opts)
}

func TestReadNamespaces(t *testing.T) {
	s := testStore()
	defer s.Close()

	_, e := s.(*store).kvStore.Set(testNamespaceKey, testNamespaces)
	require.NoError(t, e)
	nss, err := s.ReadNamespaces()
	require.NoError(t, err)
	require.NotNil(t, nss.Namespaces)
}

func TestReadNamespacesError(t *testing.T) {
	s := testStore()
	defer s.Close()

	_, e := s.(*store).kvStore.Set(testNamespaceKey, &schema.RollupRule{Uuid: "x"})
	require.NoError(t, e)
	nss, err := s.ReadNamespaces()
	require.Error(t, err)
	require.Nil(t, nss)
}

func TestReadRuleSet(t *testing.T) {
	s := testStore()
	defer s.Close()

	_, e := s.(*store).kvStore.Set(testRuleSetKey, testRuleSet)
	require.NoError(t, e)
	rs, err := s.ReadRuleSet(testNamespace)
	require.NoError(t, err)
	require.NotNil(t, rs)
}

func TestReadRuleSetError(t *testing.T) {
	s := testStore()
	defer s.Close()

	_, e := s.(*store).kvStore.Set(testRuleSetKey, &schema.Namespace{Name: "x"})
	require.NoError(t, e)
	rs, err := s.ReadRuleSet("blah")
	require.Error(t, err)
	require.Nil(t, rs)
}

func TestWriteAll(t *testing.T) {
	s := testStore()
	defer s.Close()

	rs, err := s.ReadRuleSet(testNamespaceKey)
	require.Error(t, err)
	require.Nil(t, rs)

	nss, err := s.ReadNamespaces()
	require.Error(t, err)
	require.Nil(t, nss)

	mutable := newMutableRuleSetFromSchema(t, 0, testRuleSet)
	namespaces, err := rules.NewNamespaces(0, testNamespaces)
	require.NoError(t, err)

	err = s.WriteAll(&namespaces, mutable)
	require.NoError(t, err)

	rs, err = s.ReadRuleSet(testNamespace)
	require.NoError(t, err)
	rsSchema, err := rs.ToMutableRuleSet().Schema()
	require.NoError(t, err)
	require.Equal(t, rsSchema, testRuleSet)

	nss, err = s.ReadNamespaces()
	require.NoError(t, err)
	nssSchema, err := nss.Schema()
	require.NoError(t, err)
	require.Equal(t, nssSchema, testNamespaces)
}

func TestWriteAllValidationError(t *testing.T) {
	errInvalidRuleSet := errors.New("invalid ruleset")
	v := &mockValidator{
		validateFn: func(rules.RuleSet) error { return errInvalidRuleSet },
	}
	s := testStoreWithValidator(v)
	defer s.Close()
	require.Equal(t, errInvalidRuleSet, s.WriteAll(nil, nil))
}

func TestWriteAllError(t *testing.T) {
	s := testStore()
	defer s.Close()

	rs, err := s.ReadRuleSet(testNamespaceKey)
	require.Error(t, err)
	require.Nil(t, rs)

	nss, err := s.ReadNamespaces()
	require.Error(t, err)
	require.Nil(t, nss)

	mutable := newMutableRuleSetFromSchema(t, 1, testRuleSet)
	namespaces, err := rules.NewNamespaces(0, testNamespaces)
	require.NoError(t, err)

	type dataPair struct {
		nss *rules.Namespaces
		rs  rules.MutableRuleSet
	}

	otherNss, err := rules.NewNamespaces(1, testNamespaces)
	require.NoError(t, err)

	badPairs := []dataPair{
		dataPair{nil, nil},
		dataPair{nil, mutable},
		dataPair{&namespaces, nil},
		dataPair{&otherNss, mutable},
	}

	for _, p := range badPairs {
		err = s.WriteAll(p.nss, p.rs)
		require.Error(t, err)
	}

	_, err = s.ReadRuleSet(testNamespace)
	require.Error(t, err)

	_, err = s.ReadNamespaces()
	require.Error(t, err)
}

func TestWriteRuleSetValidationError(t *testing.T) {
	errInvalidRuleSet := errors.New("invalid ruleset")
	v := &mockValidator{
		validateFn: func(rules.RuleSet) error { return errInvalidRuleSet },
	}
	s := testStoreWithValidator(v)
	defer s.Close()
	require.Equal(t, errInvalidRuleSet, s.WriteRuleSet(nil))
}

func TestWriteRuleSetError(t *testing.T) {
	s := testStore()
	defer s.Close()

	rs, err := s.ReadRuleSet(testNamespaceKey)
	require.Error(t, err)
	require.Nil(t, rs)

	nss, err := s.ReadNamespaces()
	require.Error(t, err)
	require.Nil(t, nss)

	mutable := newMutableRuleSetFromSchema(t, 1, testRuleSet)
	badRuleSets := []rules.MutableRuleSet{mutable, nil}
	for _, rs := range badRuleSets {
		err = s.WriteRuleSet(rs)
		require.Error(t, err)
	}

	err = s.WriteRuleSet(nil)
	require.Error(t, err)

	_, err = s.ReadRuleSet(testNamespace)
	require.Error(t, err)
}

func TestWriteAllNoNamespace(t *testing.T) {
	s := testStore()
	defer s.Close()

	rs, err := s.ReadRuleSet(testNamespaceKey)
	require.Error(t, err)
	require.Nil(t, rs)

	nss, err := s.ReadNamespaces()
	require.Error(t, err)
	require.Nil(t, nss)

	mutable := newMutableRuleSetFromSchema(t, 0, testRuleSet)
	namespaces, err := rules.NewNamespaces(0, testNamespaces)
	require.NoError(t, err)

	err = s.WriteAll(&namespaces, mutable)
	require.NoError(t, err)

	rs, err = s.ReadRuleSet(testNamespace)
	require.NoError(t, err)

	_, err = s.ReadNamespaces()
	require.NoError(t, err)

	err = s.WriteRuleSet(rs.ToMutableRuleSet())
	require.NoError(t, err)

	rs, err = s.ReadRuleSet(testNamespace)
	require.NoError(t, err)
	nss, err = s.ReadNamespaces()
	require.NoError(t, err)
	require.Equal(t, nss.Version(), 1)
	require.Equal(t, rs.Version(), 2)
}

func testStore() rules.Store {
	return testStoreWithValidator(nil)
}

func testStoreWithValidator(validator rules.Validator) rules.Store {
	opts := NewStoreOptions(testNamespaceKey, testRuleSetKeyFmt, validator)
	kvStore := mem.NewStore()
	return NewStore(kvStore, opts)
}

// newMutableRuleSetFromSchema creates a new MutableRuleSet from a schema object.
func newMutableRuleSetFromSchema(
	t *testing.T,
	version int,
	rs *schema.RuleSet,
) rules.MutableRuleSet {
	// Takes a blank Options stuct because none of the mutation functions need the options.
	roRuleSet, err := rules.NewRuleSetFromSchema(version, rs, rules.NewOptions())
	require.NoError(t, err)
	return roRuleSet.ToMutableRuleSet()
}

type validateFn func(rs rules.RuleSet) error

type mockValidator struct {
	validateFn validateFn
}

func (v *mockValidator) Validate(rs rules.RuleSet) error                        { return v.validateFn(rs) }
func (v *mockValidator) ValidateSnapshot(snapshot *rules.RuleSetSnapshot) error { return nil }
func (v *mockValidator) Close()                                                 {}
