package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Permission struct {
	Name         string `json:"name"`
	ResourceName string `json:"resourceName"`
	Description  string `json:"description"`
	URL          string `json:"url"`
	Action       string `json:"action"`
	Method       string `json:"method"`
}

type Role struct {
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Permissions  []string      `json:"permissions"`
	ResourcesMap []ResourceMap `json:"resourcesMap"`
}

type ResourceMap struct {
	ResourceName   string   `json:"resourceName"`
	PermissionList []string `json:"permissionList"`
}

func gen() {
	permissionsJSON, err := ioutil.ReadFile("Permissions.json")
	if err != nil {
		fmt.Println("Error reading Permissions.json:", err)
		return
	}

	rolesJSON, err := ioutil.ReadFile("Roles.json")
	if err != nil {
		fmt.Println("Error reading Roles.json:", err)
		return
	}

	var permissions []Permission
	var roles []Role

	if err := json.Unmarshal(permissionsJSON, &permissions); err != nil {
		fmt.Println("Error unmarshaling Permissions.json:", err)
		return
	}

	if err := json.Unmarshal(rolesJSON, &roles); err != nil {
		fmt.Println("Error unmarshaling Roles.json:", err)
		return
	}

	permissionsMap := make(map[string]string)
	for _, p := range permissions {
		key := fmt.Sprintf("%s,%s", p.URL, p.Method)
		permissionsMap[p.Name] = key
	}

	for i := range roles {
		var updatedPermissions []string
		for _, p := range roles[i].Permissions {
			key, exists := permissionsMap[p]
			if exists {
				updatedPermissions = append(updatedPermissions, key)
			}
		}
		roles[i].Permissions = updatedPermissions
	}

	roles2JSON, err := json.MarshalIndent(roles, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling roles2:", err)
		return
	}

	err = ioutil.WriteFile("Roles2.json", roles2JSON, 0644)
	if err != nil {
		fmt.Println("Error writing Roles2.json:", err)
		return
	}

	fmt.Println("Roles2.json generated successfully.")
}
