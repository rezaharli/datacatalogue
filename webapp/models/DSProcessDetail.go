package models

type DSProcessDetail struct {
	ID                        int
	Process_ID                DSProcesses
	Business_Term_ID          BusinessTerm
	Segment_ID                Segment
	Imm_Prec_System_ID        System
	Ultimate_Source_System_ID System
	CDE_Rationale             string
}
