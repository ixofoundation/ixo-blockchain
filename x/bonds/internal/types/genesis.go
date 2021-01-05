package types
//
//type GenesisState struct {
//	Bonds   []Bond  `json:"bonds" yaml:"bonds"`
//	Batches []Batch `json:"batches" yaml:"batches"`
//	Params  Params  `json:"params" yaml:"params"`
//}
//
//func NewGenesisState(bonds []Bond, batches []Batch, params Params) GenesisState {
//	return GenesisState{
//		Bonds:   bonds,
//		Batches: batches,
//		Params:  params,
//	}
//}
//
//func ValidateGenesis(data GenesisState) error {
//	err := ValidateParams(data.Params)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func DefaultGenesisState() GenesisState {
//	return GenesisState{
//		Bonds:   nil,
//		Batches: nil,
//		Params:  DefaultParams(),
//	}
//}
