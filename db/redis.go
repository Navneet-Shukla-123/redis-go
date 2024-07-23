package db

import "context"

func (db *Redis) SetKey(ctx context.Context, key, value string) error {
	resp := db.DB.Set(ctx, key, value, 0)
	if resp.Err() != nil {
		return resp.Err()
	}
	return nil
}

func (db *Redis) GetKey(ctx context.Context, key string) (string, error) {
	val, err := db.DB.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}
	return val, nil

}

// when side is false push to left of list
// when side is true push to right of  list
func (db *Redis) ListPush(ctx context.Context, side bool, key, value string) error {
	if !side {
		err := db.DB.LPush(ctx, key, value).Err()
		if err != nil {
			return err
		}
	} else {
		err := db.DB.RPush(ctx, key, value).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

// when side is false pop from left of list
// when side is true pop  from right of  list
func (db *Redis) ListPop(ctx context.Context, side bool, key string) (string, error) {
	if !side {
		value, err := db.DB.LPop(ctx, key).Result()
		if err != nil {
			return "", err
		} else {
			return value, nil
		}

	} else {
		value, err := db.DB.RPop(ctx, key).Result()
		if err != nil {
			return "", err
		} else {
			return value, nil
		}
	}
}
