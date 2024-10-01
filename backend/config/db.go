package config

import (
    "fmt"
    "time"
    "example.com/project/entity"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
    return db
}

/*สร้างฐานข้อมูล*/
func ConnectionDB() {
    database, err := gorm.Open(sqlite.Open("sa1.db?cache=shared"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    fmt.Println("connected database")
    db = database
}

func getDOB(year, month, day int) time.Time {
    dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
    return dob
}

func SetupDatabase() {

    db.AutoMigrate(
		
		&entity.Patient{},&entity.Employee{},&entity.Gender{},&entity.BloodType{},&entity.JobPosition{},
		
		&entity.Schedule{},&entity.Tstatus{},&entity.Treatment{},
		
		&entity.DentalRecord{},&entity.Status{},
		
		&entity.Payment{},&entity.PaymentMethod{},
		
		&entity.Equipments{}, &entity.Requisitions{}, &entity.Restocks{},
	)
	

	// Treatment
    TreatmentTeethExamination := entity.Treatment{TreatmentName: "ตรวจฟัน"}
    TreatmentCleaning := entity.Treatment{TreatmentName: "ขูดหินปูน"}
	TreatmentFillTeeth := entity.Treatment{TreatmentName: "อุดฟัน"}
	TreatmentPullTooth := entity.Treatment{TreatmentName: "ถอนฟัน"}
	TreatmentToothImpacdtionRemoval := entity.Treatment{TreatmentName: "ผ่าฟันคุด"}
	TreatmentRootCanalTherapy := entity.Treatment{TreatmentName: "รักษารากฟัน"}
	TreatmentCrown := entity.Treatment{TreatmentName: "ครอบฟัน"}
	TreatmentFluorideApplication := entity.Treatment{TreatmentName: "เคลือบฟลูออไรด์"}
	TreatmentOrthodontics := entity.Treatment{TreatmentName: "จัดฟัน"}

	db.FirstOrCreate(&TreatmentTeethExamination, &entity.Treatment{TreatmentName: "ตรวจฟัน"})
    db.FirstOrCreate(&TreatmentCleaning, &entity.Treatment{TreatmentName: "ขูดหินปูน"})
    db.FirstOrCreate(&TreatmentFillTeeth, &entity.Treatment{TreatmentName: "อุดฟัน"})
    db.FirstOrCreate(&TreatmentPullTooth, &entity.Treatment{TreatmentName: "ถอนฟัน"})
	db.FirstOrCreate(&TreatmentRootCanalTherapy, &entity.Treatment{TreatmentName: "ผ่าฟันคุด"})
	db.FirstOrCreate(&TreatmentToothImpacdtionRemoval, &entity.Treatment{TreatmentName: "รักษารากฟัน"})
	db.FirstOrCreate(&TreatmentCrown, &entity.Treatment{TreatmentName: "ครอบฟัน"})
	db.FirstOrCreate(&TreatmentFluorideApplication, &entity.Treatment{TreatmentName: "เคลือบฟลูออไรด์"})
	db.FirstOrCreate(&TreatmentOrthodontics, &entity.Treatment{TreatmentName: "จัดฟัน"})

	// TStatus
	TStatusPending := entity.Tstatus{TStatusName: "รอดำเนินการ"}
	TStatusDone := entity.Tstatus{TStatusName: "สำเร็จ"}
	TStatusCancel := entity.Tstatus{TStatusName: "ยกเลิก"}

	db.FirstOrCreate(&TStatusPending, &entity.Tstatus{TStatusName: "รอดำเนินการ"})
	db.FirstOrCreate(&TStatusDone, &entity.Tstatus{TStatusName: "สำเร็จ"})
	db.FirstOrCreate(&TStatusCancel, &entity.Tstatus{TStatusName: "ยกเลิก"})
	// Gender
	GenderMale := entity.Gender{Sex : "ชาย"}
	GenderFemale := entity.Gender{Sex : "หญิง"}

	db.FirstOrCreate(&GenderMale, &entity.Gender{Sex : "ชาย"})
	db.FirstOrCreate(&GenderFemale, &entity.Gender{Sex : "หญิง"})

	// BloodType
	BloodO:= entity.BloodType{BloodGroup: "O"}
	BloodA:= entity.BloodType{BloodGroup: "A"}
	BloodB:= entity.BloodType{BloodGroup: "B"}
	BloodAB:= entity.BloodType{BloodGroup: "AB"}

	db.FirstOrCreate(&BloodO, &entity.BloodType{BloodGroup : "O"})
	db.FirstOrCreate(&BloodA, &entity.BloodType{BloodGroup : "A"})
	db.FirstOrCreate(&BloodB, &entity.BloodType{BloodGroup : "B"})
	db.FirstOrCreate(&BloodAB, &entity.BloodType{BloodGroup : "AB"})

	//แผนกพนักงาน
	JobPositionDentist := entity.JobPosition{Job: "ทันตแพทย์"}
	JobPositionPatientService := entity.JobPosition{Job: "เจ้าหน้าที่บริการคนไข้"}
	JobPositionAdmin := entity.JobPosition{Job: "ผู้ดูแลระบบ"}

	db.FirstOrCreate(&JobPositionDentist, &entity.JobPosition{Job: "ทันตแพทย์"})
	db.FirstOrCreate(&JobPositionPatientService, &entity.JobPosition{Job: "เจ้าหน้าที่บริการคนไข้"})
	db.FirstOrCreate(&JobPositionAdmin, &entity.JobPosition{Job: "ผู้ดูแลระบบ"})

	//วิธีชำระเงิน
	transfer:= entity.PaymentMethod{MethodName: "โอน"}//เอาไว้เช็คในฐานข้อมูลว่ามีคำนี้หรือยัง
	cash := entity.PaymentMethod{MethodName: "เงินสด"}
	creditcard := entity.PaymentMethod{MethodName: "บัตรเครดิต"}
	db.FirstOrCreate(&transfer, &entity.PaymentMethod{MethodName: "โอน"})
	db.FirstOrCreate(&cash, &entity.PaymentMethod{MethodName: "เงินสด"})
	db.FirstOrCreate(&creditcard, &entity.PaymentMethod{MethodName: "บัตรเครดิต"})

	//status
	StatusPaid:= entity.Status{StatusName: "ชำระแล้ว"}
	StatusNotPaid:= entity.Status{StatusName: "ยังไม่ชำระ"}
	db.FirstOrCreate(&StatusPaid, &entity.Status{StatusName : "ชำระแล้ว"})
	db.FirstOrCreate(&StatusNotPaid, &entity.Status{StatusName: "ยังไม่ชำระ"})


	dob := getDOB(2011, 4, 2)
	dob2 := getDOB(2000, 2, 1)
	dob3 := getDOB(1999, 6, 6)
	
	Patient3:=entity.Patient{
		FirstName :		"สมหญิง",
		LastName :		"สุขขี",
		Birthday :		dob,
		Weight : 		45,
		Height : 		150,
		DrugAllergy :		"ความดัน",
		Chronicdisease : 	"ย่าฆ่าเชื้อ",
		Tel :"0610000000",
		BloodTypeID :1,
		GenderID:1,
		}
	db.FirstOrCreate(&Patient3)
	
	//
	Patient := &entity.Patient{
		FirstName: 		"นรชาติ",
		LastName:  		"ติวางวาย",
		Birthday:   	dob2,
		Weight:   		66,
		Height:  		166,
		GenderID:		1,
		BloodTypeID:	1,
		DrugAllergy:	"-",
		Chronicdisease:	"-",
		Tel:			"0000000000",
	}
	db.FirstOrCreate(&Patient, entity.Patient{
		FirstName: "นรชาติ",
		LastName:  "ติวางวาย",
	})

	Patient2 := &entity.Patient{
		FirstName: 		"ธนภูมิ",
		LastName:  		"กินอิ่ม",
		Birthday:   	dob3,
		Weight:   		66,
		Height:  		176,
		GenderID:		1,
		BloodTypeID:	2,
		DrugAllergy:	"-",
		Chronicdisease:	"-",
		Tel:			"1111111111",
	}
	db.FirstOrCreate(&Patient2, entity.Patient{
		FirstName: "ธนภูมิ",
		LastName:  "กินอิ่ม",
	})
	
	
	
	
	
	//
	hashedPassword, _ := HashPassword("123456")
	BirthDay, _ := time.Parse("2006-01-02", "1988-11-12")

	Employee1	:=	&entity.Employee{
		FirstName: 		"รามณรงค์",
       	LastName:		"พันธเดช",
		Birthday:		BirthDay,
		Address:		"ประเทศไทย",
		Tel:			"0822222222",
       	Email:     		"admin@gmail.com",
       	Password:  		hashedPassword,
       	GenderID:  		1,
		JobPositionID:  3,
	}

	db.FirstOrCreate(Employee1, &entity.Employee{
		Email: 			"admin@gmail.com",
	})


	hashedPassword2, _ := HashPassword("123456")
	

	Employee2	:=	&entity.Employee{
		FirstName: 		"สมชาย",
       	LastName:		"ใจดี",
		Birthday:		BirthDay,
		Address:		"ประเทศไทย",
		Tel:			"0811111111",
       	Email:     		"sa2@gmail.com",
       	Password:  		hashedPassword2,
       	GenderID:  		1,
		JobPositionID:  2,
	}

	db.FirstOrCreate(Employee2, &entity.Employee{
		Email: 			"sa2@gmail.com",
	})

	

	// payment
	//hashedPassword2, _ := HashPassword("123456")
	//BirthDay2, _ := time.Parse("2006-01-02T00:00:00Z", "1988-11-12T00:00:00Z")
	//พนักงาน
	Employee := &entity.Employee{
 
		FirstName: "เนตรนภา",
 
		LastName:  "สารวัน",
 
		Email:     "gift@gmail.com",
 
		Address: "โคราช",
 
		Password: hashedPassword,
 
		Birthday:  BirthDay,
 
		GenderID:  2,

		JobPositionID: 2,
		
		Tel: "0987654321",
 
	}

	EmployeeDoctor := &entity.Employee{
 
		FirstName: "สมชาย",
 
		LastName:  "สายสุนทรีย์",
 
		Email:     "doctor@gmail.com",
 
		Address: "โคราช",
 
		Password: hashedPassword,
 
		Birthday:  BirthDay,
 
		GenderID:  2,

		JobPositionID: 1,
		
		Tel: "0631114444",
 
	}
 
	db.FirstOrCreate(Employee, &entity.Employee{
 
		Email: "gift@gmail.com",
 
	})
	db.FirstOrCreate(EmployeeDoctor, &entity.Employee{
 
		Email: "doctor@gmail.com",
 
	})


	//Treatment
	ScrapeWayTartar:= entity.Treatment{TreatmentName: "ขูดหินปูน" }
	db.FirstOrCreate(&ScrapeWayTartar, &entity.Treatment{TreatmentName: "ขูดหินปูน"})
	//คนไข้
	Patient23:=entity.Patient{
		FirstName :"สมหญิง",
		LastName :"สุขขี",
		Birthday :BirthDay,
		Weight : 45,
		Height : 150,
		DrugAllergy :"ความดัน",
		Chronicdisease : "ย่าฆ่าเชื้อ",
		Tel :"061-000-0000",
		BloodTypeID :1,
		GenderID:1,
		}
	db.FirstOrCreate(&Patient23)

	
	
	//ชำระเงิน
	NowDate := time.Now()//เวลาปัจจุบัน
	Payment:=entity.Payment{
		Date :NowDate,

		PaymentMethodID : 1,

		EmployeeID : 2,
	}
	db.FirstOrCreate(&Payment)

	record := entity.DentalRecord{
		Date :NowDate,
		Description :"ฟันพุเยอะมาก",
		Fees :500.00,
		Installment: 0,
		NumberOfInstallment: "0/0",

		PatientID :1,
		
		EmployeeID :3,

		TreatmentID :4,
		
		StatusID :2,

		PaymentID: nil,
		}
	db.FirstOrCreate(&record,entity.DentalRecord{
		Description :"ฟันพุเยอะมาก",
	})


	record2 := entity.DentalRecord{
		Date :NowDate,
		Description :"จัดฟันครั้งแรก",
		Fees :400.00,
		Installment: 40000.00,
		NumberOfInstallment: "1/12",

		PatientID :2,
		
		EmployeeID :3,

		TreatmentID :9,
		
		StatusID :2,

		PaymentID: nil,
		}
	db.FirstOrCreate(&record2,entity.DentalRecord{
		Description :"จัดฟันครั้งแรก",
	})

	record3 := entity.DentalRecord{
		Date :NowDate,
		Description :"ตรวจช่องปากหลายจุด",
		Fees :1200.00,
		Installment: 0,
		NumberOfInstallment: "0/0",

		PatientID :3,
		
		EmployeeID :3,

		TreatmentID :1,
		
		StatusID :2,

		PaymentID: nil,
		}
	db.FirstOrCreate(&record3,entity.DentalRecord{
		Description :"ตรวจช่องปากหลายจุด",
	})




		//อุปกรณ์
		equipment1 := entity.Equipments{
			EquipmentName: "Orthodontic Wires Steel (ลวดจัดฟัน)",
			Unit:          "เส้น",
			Cost:          400.66,
			Quantity:      96,
		}
		db.FirstOrCreate(&equipment1, entity.Equipments{
			EquipmentName: "Orthodontic Wires Steel (ลวดจัดฟัน)",
		})
	
		equipment2 := entity.Equipments{
			EquipmentName: "Cheek Retractors (แผ่นกันลิ้น)",
			Unit:          "แพ็ค",
			Cost:          450.56,
			Quantity:      80,
		}
		db.FirstOrCreate(&equipment2, entity.Equipments{
			EquipmentName: "Cheek Retractors (แผ่นกันลิ้น)",
		})

		equipment3 := entity.Equipments{
			EquipmentName: "Acrylic (ครอบฟันชั่วคราว)",
			Unit:          "ชุด",
			Cost:          2900.99,
			Quantity:      266,
		}
		db.FirstOrCreate(&equipment3, entity.Equipments{
			EquipmentName: "Acrylic (ครอบฟันชั่วคราว)",
		})
		
		equipment4 := entity.Equipments{
			EquipmentName: "Orthodontic Wires Nickel-Titanium (ลวดจัดฟัน)",
			Unit:          "เส้น",
			Cost:          550.50,
			Quantity:      310,
		}
		db.FirstOrCreate(&equipment4, entity.Equipments{
			EquipmentName: "Orthodontic Wires Nickel-Titanium (ลวดจัดฟัน)",
		})
		
		equipment5 := entity.Equipments{
			EquipmentName: "Orthodontic Wires Beta-Titanium (ลวดจัดฟัน)",
			Unit:          "เส้น",
			Cost:          600.50,
			Quantity:      250,
		}
		db.FirstOrCreate(&equipment5, entity.Equipments{
			EquipmentName: "Orthodontic Wires Beta-Titanium (ลวดจัดฟัน)",
		})

		equipment6 := entity.Equipments{
			EquipmentName: "X-ray Film (ฟิล์มเอ็กซเรย์)",
			Unit:          "เเผ่น",
			Cost:          100.60,
			Quantity:      349,
		}
		db.FirstOrCreate(&equipment6, entity.Equipments{
			EquipmentName: "X-ray Film (ฟิล์มเอ็กซเรย์)",
		})

		equipment7 := entity.Equipments{
			EquipmentName: "Dental Polishing Strips (แผ่นขัดฟัน)",
			Unit:          "เเผ่น",
			Cost:          150.80,
			Quantity:      340,
		}
		db.FirstOrCreate(&equipment7, entity.Equipments{
			EquipmentName: "Dental Polishing Strips (แผ่นขัดฟัน)",
		})



	
		//เบิก
		// แปลงสตริงของวันที่และเวลาให้เป็น time.Time
		customTime1, _ := time.Parse("2006-01-02 15:04:05", "2024-09-26 19:00:01")
		customTime2, _ := time.Parse("2006-01-02 15:04:05", "2024-09-26 19:00:20")
		customTime3, _ := time.Parse("2006-01-02 15:04:05", "2024-09-27 19:00:09")
		customTime4, _ := time.Parse("2006-01-02 15:04:05", "2024-09-27 19:00:29")
	
		requisition1 := entity.Requisitions{
			RequisitionQuantity: 7,
			Time:                customTime1,
			Note:                "ใช้เปลี่ยน",
			EquipmentID:         2,
			EmployeeID:          2,
		}
		db.FirstOrCreate(&requisition1, entity.Requisitions{
			Time: customTime1,
		})
	
		requisition2 := entity.Requisitions{
			RequisitionQuantity: 6,
			Time:                customTime2,
			Note:                "-",
			EquipmentID:         1,
			EmployeeID:          3,
		}
		db.FirstOrCreate(&requisition2, entity.Requisitions{
			Time: customTime2,
		})
		
		requisition3 := entity.Requisitions{
			RequisitionQuantity: 10,
			Time:                customTime3,
			Note:                "-",
			EquipmentID:         6,
			EmployeeID:          2,
		}
		db.FirstOrCreate(&requisition3, entity.Requisitions{
			Time: customTime3,
		})

		requisition4 := entity.Requisitions{
			RequisitionQuantity: 15,
			Time:                customTime4,
			Note:                "-",
			EquipmentID:         6,
			EmployeeID:          2,
		}
		db.FirstOrCreate(&requisition4, entity.Requisitions{
			Time: customTime4,
		})


		//เติม
		customTime01, _ := time.Parse("2006-01-02 15:04:05", "2024-08-01 08:40:10")
		customTime02, _ := time.Parse("2006-01-02 15:04:05", "2024-08-01 08:45:39")
		customTime03, _ := time.Parse("2006-01-02 15:04:05", "2024-08-15 08:30:09")
		customTime04, _ := time.Parse("2006-01-02 15:04:05", "2024-08-15 08:30:29")
	
		restock01 := entity.Restocks{
			RestockQuantity: 300,
			ReceivingDate:   customTime01,
			EquipmentID:     1,
			EmployeeID:      3,
		}
		db.FirstOrCreate(&restock01, entity.Restocks{
			ReceivingDate: customTime01,
		})
	
		restock02 := entity.Restocks{
			RestockQuantity: 400,
			ReceivingDate:   customTime02,
			EquipmentID:     2,
			EmployeeID:      2,
		}
		db.FirstOrCreate(&restock02, entity.Restocks{
			ReceivingDate: customTime02,
		})

		restock03 := entity.Restocks{
			RestockQuantity: 400,
			ReceivingDate:   customTime03,
			EquipmentID:     6,
			EmployeeID:      2,
		}
		db.FirstOrCreate(&restock03, entity.Restocks{
			ReceivingDate: customTime03,
		})

		restock04 := entity.Restocks{
			RestockQuantity: 300,
			ReceivingDate:   customTime04,
			EquipmentID:     4,
			EmployeeID:      3,
		}
		db.FirstOrCreate(&restock04, entity.Restocks{
			ReceivingDate: customTime04,
		})




		record1 := entity.DentalRecord{
			Date :NowDate,
			Description :"ทำการถอนฟันซี่ที่ 18 (ฟันกรามล่างขวา) เนื่องจากฟันผุมากและไม่สามารถรักษาได้ แนะนำการดูแลหลังถอนฟันให้คนไข้ งดใช้ฟันฝั่งที่ถอน หลีกเลี่ยงอาหารร้อนหรือแข็ง",
			Fees :1500.00,
			Installment: 0,
			NumberOfInstallment: "0/0",
	
			PatientID :2,
			
			EmployeeID :1,
	
			TreatmentID :4,
			
			StatusID :2,
	
			PaymentID: nil,
			}
		db.FirstOrCreate(&record1)
}
