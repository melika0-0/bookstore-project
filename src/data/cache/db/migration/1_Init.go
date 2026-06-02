package migration

import ()
var logger = logging.NewLogger(config.GetConfig)
func Up_1 () {
	database := db.GetDb()
	tables := []interface{}{}

}

func Down_1 (){
	
}