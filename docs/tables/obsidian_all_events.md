# Table: obsidian_all_events

This table shows data for Obsidian All Events.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|actors|`json`|
|targets|`json`|
|browser_id|`utf8`|
|datetime|`timestamp[us, tz=UTC]`|
|device_id|`utf8`|
|event_frequency|`utf8`|
|human_readable_description|`json`|
|id (PK)|`utf8`|
|ip_address|`json`|
|is_privileged_event|`bool`|
|location|`json`|
|obsidian_event_name|`json`|
|service|`json`|
|session_id|`utf8`|
|tenant_id|`utf8`|
|tenant_object|`json`|
|sub_service|`utf8`|
|status|`utf8`|
|user_agent|`json`|