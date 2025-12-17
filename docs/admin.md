## Introduction
This document describes how administrators can manage the DoData platform, including user management, data oversight, and system configuration.

## Contents
- [Introduction](#Introduction)
- [Pocketbase](#pocketbase)
  - [Users](#users)
  - [Rules](#rules)
    - [Example](#example)
  - [Recommendations](#recommendations)
  - [Export](#export)

### Pocketbase
Pocketbase is used as the backend for the DoData platform. Administrators can access the Pocketbase admin dashboard to manage users, data collections, and system settings.

:warning:  
Pocketbase administrators can modify or delete critical data collections. This can cause the system to malfunction.  
To avoid this, records are only added, modified or deleted for the following collections:
- `users`
- `rules`
- `recommendation_templates`

When managing records in a colleciton the following primary actions can be performed: 
- To create a new record, click the **New Record** button and fill in the required fields.
- To edit a record, click on the record and modify the fields as necessary.
- To delete a record, select the record and click the ellipsis (...) button in the top right corner, then choose **Delete**.

#### Users
Users can be managed through the `users` collection. Here administrators can add, edit, or remove users as needed.
Users can also register themselves through the frontend interface, but administrators have the ability to manage all user accounts here.
<img src="/docs/users.png" />

#### Rules
Rules for generating insights and recommendations can be managed through the `rules` collection. Administrators can create new rules, modify existing ones, or delete rules that are no longer needed.
Rules are defined by four properties:
- `situation`: This along with the `threshold` form the specific condition that triggers the rule.
    - Currently three different situations are supported:
        - `trending_up`: Triggered when the data from one year shows an **increasing** trend exceeding the threshold compared to the previous year.
        - `trending_down`: Triggered when the data from one year shows an **decreasing** trend exceeding the threshold compared to the previous year.
        - `trend_steady`: Triggered when the data from one year shows a **stable** trend within the threshold compared to the previous year.
- `domain`: The domain the rule applies, either: `El`, `Varme` or `Gas`.
- `type`: The type of data, either: `CO2 Emissions`, `Energy Consumption Renewable` or `Energy Consumption Non-Renewable`.
- `organization_category`: The category of the organization, from a predefined set e.g., `Produktionsvirksomhed`, `Servicevirksomhed`, or `Handelsvirksomhed`.
- `threshold`: A numeric value representing the percentage change that must be exceeded for the rule to trigger.

Rules are triggered when new data is imported that meets the conditions defined by the rule.
If a condition has not been set it will apply to all values for that property. E.g. if no `domain` has been selected the rule will apply to all domains.

##### Example  
A rule which triggers when the CO2 emissions for heating, for any company (no `organization_category`) increases beyond 20% (`threshold` set to 20) compared to the previous year would be defined as:
<img src="/docs/rule.png" />

#### Recommendations
The solution automatically generates recommendations based on the rules defined in the `rules` collection.
However to display a specific recommendation widget when a rule is triggered, administrators must define recommendation templates in the `recommendation_templates` collection.

The primary fields for a recommendation template are:
- `rule`: A relation field linking to the specific rule that this template corresponds to.
- `title`: The title of the recommendation displayed to the user.
- `text`: The detailed recommendation text providing insights and suggestions to the user.

<img src="/docs/recommendation.png" height=400 />


#### Export
In order for export to be accessible to municipality users, the user must be assigned the `MUNICIPALITY` role in Pocketbase.

Furthermore the export functionality relies on the environment variable `MUNICIPALITY_EXPORT` being set to the minimum number of organizations required to enable export. 

E.g. if `MUNICIPALITY_EXPORT=5`, the export functionality will only be available to municipality users if there are at least 5 organizations which have registered and imported data.

The variable should be set as part of the Pocketbase hosting configuration.
