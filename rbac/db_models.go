package rbac

type ui_Item struct {
	Role_Id     string `db:"role_id"`
	Id          int    `db:"ui_id"`
	Key         string `db:"ui_key"`
	UiType      string `db:"ui_type"`
	Description string `db:"description"`
	Parent_id   int    `db:"parent_ui_id"`
	IsAllow     bool   `db:"isallow"`
}
