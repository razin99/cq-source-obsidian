# Table: obsidian_global_posture_rules

This table shows data for Obsidian Global Posture Rules.

The primary key for this table is **rule_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|rule_id (PK)|`utf8`|
|platform_id|`utf8`|
|product_ids|`list<item: utf8, nullable>`|
|security_domain|`json`|
|name|`utf8`|
|risk_level|`utf8`|
|standard_ids|`list<item: utf8, nullable>`|
|control_ids|`list<item: utf8, nullable>`|
|obsidian_rule|`bool`|
|release_label|`utf8`|
|default_risk_level|`utf8`|
|benchmark_value_type|`utf8`|
|rule_value_type|`utf8`|
|description|`utf8`|
|remediation_instructions|`utf8`|
|exceptions_count|`json`|
|tenant_states|`json`|
|total_violations|`int64`|
|correction_score_change|`float64`|