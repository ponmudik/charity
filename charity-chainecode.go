package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Goal struct {
	//GoalId          string `json:"goalId"`
	CharityId       string `json:"charityId"`
	CharityName     string `json:"charityName"`
	EstimatedAmount string `json:"estimatedAmount"`
	Type            string `json:"type"`
	Details         string `json:"details"`
	Timestamp       string `json:"timestamp"`
}

type Donation struct {
	//DonationId string `json:"donationId"`
	UserId    string `json:"userId"`
	GoalId    string `json:"goalId"`
	CharityId string `json:"charityId"`
	Amount    string `json:"amount"`
	Timestamp string `json:"timestamp"`
}

type Expenditure struct {
	//ExpenditureId string `json: "expenditureId"`
	GoalId    string `json:"goalId"`
	CharityId string `json:"charityId"`
	SpentBy   string `json:"spentBy"`
	Purpose   string `json:"purpose"`
	Amount    string `json:"amount"`
	Timestamp string `json:"timestamp"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "publishGoal" {
		return s.publishGoal(APIstub, args)
	} else if function == "donateFund" {
		return s.donateFund(APIstub, args)
	} else if function == "addExpenditure" {
		return s.addExpenditure(APIstub, args)
	} else if function == "queryGoal" {
		return s.queryGoal(APIstub, args)
	} else if function == "queryDonation" {
		return s.queryDonation(APIstub, args)
	} else if function == "queryExpenditure" {
		return s.queryExpenditure(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) publishGoal(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var goal = Goal{CharityId: args[1], CharityName: args[2], EstimatedAmount: args[3], Type: args[4], Details: args[5], Timestamp: args[6]}

	goalAsBytes, _ := json.Marshal(goal)
	err := APIstub.PutState(args[0], goalAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to add charity goals: %s", args[0]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) donateFund(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var donation = Donation{UserId: args[1], GoalId: args[2], CharityId: args[3], Amount: args[4], Timestamp: args[5]}

	donationAsBytes, _ := json.Marshal(donation)
	err := APIstub.PutState(args[0], donationAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to donate funds: %s", args[0]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) addExpenditure(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var expenditure = Expenditure{GoalId: args[1], CharityId: args[2], SpentBy: args[3], Purpose: args[4],Amount: args[5] Timestamp: args[6]}

	expenditureAsBytes, _ := json.Marshal(expenditure)
	err := APIstub.PutState(args[0], expenditureAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to add goal expenditure: %s", args[0]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) queryGoal(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	goalAsBytes, _ := APIstub.GetState(args[0])
	if goalAsBytes == nil {
		return shim.Error("Could not find goal")
	}
	return shim.Success(goalAsBytes)
}

func (s *SmartContract) queryDonation(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	donationAsBytes, _ := APIstub.GetState(args[0])
	if donationAsBytes == nil {
		return shim.Error("Could not find donation")
	}
	return shim.Success(donationAsBytes)
}

func (s *SmartContract) queryExpenditure(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	expentidureAsBytes, _ := APIstub.GetState(args[0])
	if expentidureAsBytes == nil {
		return shim.Error("Could not find expenditure")
	}
	return shim.Success(expentidureAsBytes)
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
