// Package model provides entities used in the business intelligence service
package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User includes all account information
type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // mongoDB requires primitive.ObjectID
	Name       string             `json:"name"`                               // employee name
	Username   string             `json:"username"`                           // username with no restrictions
	Password   string             `json:"password"`                           // store as hash
	Birthday   time.Time          `json:"birthday"`                           // birthdate
	Email      string             `json:"email"`                              // email
	Earnings   []Earning          `json:"earnings"`                           // complete list of earnings Note: Use Filter
	Deductions []Deduction        `json:"deductions"`                         // complete list of deductions Note: Use Filter
	Leaves     []Leave            `json:"leaves"`                             // complete list of leaves Note: Use Filter
	Payslips   []Payslip          `json:"payslip"`                            // Generated during payroll
}

// Payslip includes snapshot of a payday's total earnings and deductions
type Payslip struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // mongoDB requires primitive.ObjectID
	DateStart time.Time          `json:"dateStart"`                          // earning computation start
	DateEnd   time.Time          `json:"dateEnd"`                            // earning computation end
	Earning   Earning            `json:"earnings"`                           // payroll period
	Deduction Deduction          `json:"deductions"`                         // payroll deductions
	Leave     Leave              `json:"leaves"`                             // payroll period leaves
}

// Earning includes all earnings regardless if taxable only
type Earning struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // mongoDB requires primitive.ObjectID
	EarningUpdateDate time.Time          `json:"earningUpdateDate"`                  // computation follows latest earning UpdateDate
	Rate              float32            `json:"rate"`                               // computed from BasicSalary OR fixed
	BasicSalary       float32            `json:"basicSalary"`                        // arbitrary value
	Deminimis         float32            `json:"deminimis"`                          // arbitrary value
	Allowances        []Allowance        `json:"allowances"`                         // arbitrary list
	Holiday           float32            `json:"holiday"`                            // get from holiday list, compute for payslip duration
	Birthday          float32            `json:"birthday"`                           // gets arbitrary gift
	Overtime          float32            `json:"overtime"`                           // computes from employee overtime list
	SlVl              float32            `json:"SlVl"`                               // computes from leaves list
	Adjustment        float32            `json:"adjustment"`                         // arbitrary value
	TotalEarnings     float32            `json:"totalEarnings"`                      // computes from all of the above
}

// Leave includes all types such as sick leave, vacation leave, absence over leave
type Leave struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // mongoDB requires primitive.ObjectID
	Type         string             `json:"type"`                               // sick leave, vacation leave, absence over leave
	Reason       string             `json:"reason"`                             // reason for leave
	FilingDate   time.Time          `json:"filingDate"`                         // employee leave reason
	ApprovalDate time.Time          `json:"approvalDate"`                       // hr approval
	LeaveDate    []time.Time        `json:"leaveDate"`                          // actual leave dates
	Duration     float32            `json:"duration"`                           //  0~1 e.g. 0.5 halfday, 1 wholeday
}

type Allowance struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // mongoDB requires primitive.ObjectID
	Name   string             `json:"name"`                               // allowance can be ecola, transpo, meals, others
	Amount float32            `json:"amount"`                             // arbitrary input by managing officers
}

// Deduction includes all deductions regardless if internal or not
type Deduction struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"` // mongoDB requires primitive.ObjectID
	DateStart       time.Time          `json:"dateStart"`                          // deduction computation start
	DateEnd         time.Time          `json:"dateEnd"`                            // deduction computation end
	TotalDeductions float32            `json:"totalDeductions"`                    // computes from all of the above
}
