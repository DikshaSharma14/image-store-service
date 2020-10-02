package sim_sqlite

import (
	"image-store-service/constants"
	. "image-store-service/logger"
	"reflect"
)

// Creates a row in the table mapped to the parameter interface
func Post(data interface{}) error {
	Logger.Trace(constants.LogBeginFunc)

	SqliteInstance.AutoMigrate(data)
	Logger.Debug("Persisting new record : ", &data)
	result := SqliteInstance.Create(data)

	err := result.Error
	if err != nil {
		Logger.Warn(err.Error(), err)
		return result.Error
	}

	Logger.Trace(constants.LogFinishFunc)
	return nil
}

// Finds the first record with the specified ID from the tables mapped to the parameter interface.
// Returns err if table not found
func Get(id string, data interface{}) error {
	Logger.Trace(constants.LogBeginFunc)

	Logger.Debug("Getting record with id : ", id)
	result := SqliteInstance.First(data, "ID = ?", id)

	err := result.Error
	if err != nil {
		Logger.Warn(err.Error(), err)
		return result.Error
	}

	Logger.Trace(constants.LogFinishFunc)
	return nil
}

// Updates the row in the table mapped to the parameter interface.
func Update(data interface{}) error {
	Logger.Trace(constants.LogBeginFunc)

	Logger.Debug("Updating record : ", data)
	result := SqliteInstance.Save(data)

	err := result.Error
	if err != nil {
		Logger.Warn(err.Error(), err)
		return result.Error
	}

	Logger.Trace(constants.LogFinishFunc)
	return nil
}

// Deletes the first record with the specified ID from the tables mapped to the parameter interface.
func Delete(id string, data interface{}) error {
	Logger.Trace(constants.LogBeginFunc)

	Logger.Debug("Deleting record with id : ", id)
	result := SqliteInstance.Delete(data, "ID = ? ", id)

	err := result.Error
	if err != nil {
		Logger.Warn(err.Error(), err)
		return result.Error
	}

	Logger.Trace(constants.LogFinishFunc)
	return nil
}

// Lists all rows from the table mapped to the parameter interface
func List(data interface{}) error {
	Logger.Trace(constants.LogBeginFunc)

	Logger.Debug("Listing all records of type : ", reflect.TypeOf(data))
	result := SqliteInstance.Find(data)

	err := result.Error
	if err != nil {
		Logger.Warn(err.Error(), err)
		return result.Error
	}

	Logger.Trace(constants.LogFinishFunc)
	return nil
}

// Lists all rows from the table mapped to the parameter interface based on the condition
func ListBasedOnQuery(data interface{}, query string, args ...interface{}) error {
	Logger.Trace(constants.LogBeginFunc)

	Logger.Debug("Listing all records of type : ", reflect.TypeOf(data))
	result := SqliteInstance.Where(query, args...).Find(data)

	err := result.Error
	if err != nil {
		Logger.Warn(err.Error(), err)
		return result.Error
	}

	Logger.Trace(constants.LogFinishFunc)
	return nil
}

// Lists all rows from the table mapped to the parameter interface based on the condition
func ListBasedOnCondition(data interface{}, condition interface{}) error {
	Logger.Trace(constants.LogBeginFunc)

	Logger.Debug("Listing all records of type : ", reflect.TypeOf(data))
	result := SqliteInstance.Where(condition).Find(data)

	err := result.Error
	if err != nil {
		Logger.Warn(err.Error(), err)
		return result.Error
	}

	Logger.Trace(constants.LogFinishFunc)
	return nil
}

func DeleteBasedOnCondition(data interface{}, condition interface{}) error {
	Logger.Trace(constants.LogBeginFunc)
	result := SqliteInstance.Delete(data, condition)
	err := result.Error
	if err != nil {
		Logger.Warn(err.Error(), err)
		return result.Error
	}

	Logger.Trace(constants.LogFinishFunc)
	return nil
}
