package exec_stat_data

// 用户转化ID
var ZhiDunBusinessUser = ExecData{
	Data: `
889073384509596
177988260615621
`,
	Tmpl: `
SELECT 
  t.ext_user_id,
  bui.user_id AS user_id
FROM (
  {{- range $row := .rows }}
  SELECT '{{ index $row 1 }}' AS ext_user_id
  UNION ALL
  {{- end }}
) AS t
LEFT JOIN business_users_info_1340 bui 
  ON bui.ext_user_id = t.ext_user_id;
`,
}
