// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package pump

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Creates the global state.
type Initialize struct {

	// [0] = [WRITE] global
	//
	// [1] = [WRITE, SIGNER] user
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	nd := &Initialize{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetGlobalAccount sets the "global" account.
func (inst *Initialize) SetGlobalAccount(global ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(global).WRITE()
	return inst
}

// GetGlobalAccount gets the "global" account.
func (inst *Initialize) GetGlobalAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUserAccount sets the "user" account.
func (inst *Initialize) SetUserAccount(user ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(user).WRITE().SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *Initialize) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Initialize) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Initialize) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst Initialize) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Initialize,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Initialize) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Global is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *Initialize) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Initialize")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       global", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         user", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj Initialize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *Initialize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewInitializeInstruction declares a new Initialize instruction with the provided parameters and accounts.
func NewInitializeInstruction(
	// Accounts:
	global ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *Initialize {
	return NewInitializeInstructionBuilder().
		SetGlobalAccount(global).
		SetUserAccount(user).
		SetSystemProgramAccount(systemProgram)
}
