## Introduction
This document describes how end users can use the DoData platform to manage and visualize their energy consumption data.

## Contents
- [Introduction](#Introduction)
- [Creating a user](#creating-a-user)
- [Managing data](#managing-data)
- [Dashboard](#dashboard)
- [Exporting data](#exporting-data)

## Creating a user
The create user flow can be accessed from the main landing page by clicking the **Ny bruger** link, beneath the login form.

<img src="/docs/create_user.png" height="700" />

Two types of users can be created:
- `Virksomhed`: This user type is intended for companies looking to monitor and analyze their data.
- `Kommune`: This user type is intended for municipalities that want to export aggregated data from the system.

:exclamation:
`Kommune` users must be approved by an administrator before they can access the system.
This is done using the Pocketbase backend see the [administrator documentation](/docs/admin.md) for more information.

## Managing data
To manage your data navigate to the **Data** tab in the top menu.

Here your can upload new data using the button **Indlæs Excel-fil fra Klilmakompasset**.
In order for the system to correctly process your data there are a few requirements that the data file must comply with:

1. It must be a `.xlsx` file exported from Klimakompasset.
2. The data in the file must span a whole year. *To confirm this check the dates on the second worksheet in the exported file labeled: `Stamdata`*  

<img src="/docs/upload_data.png" />

To delete your data use the **Slet** button on each yearly import to delete all the data for that year.

## Dashboard
On the **Dashboard** tab you can view visualizations of your data.  

If the system has detected one or more insights from your data these will be displayed at the top portion of the dashboard, you can navigate between them using the **Forrige** and **Næste** buttons.

Below you will see visualizations of your data. Data is grouped by year and the visualizations show data displaying the distribution of your energy consumption with various comparison options.

<img src="/docs/dashboard.png" />

### Exporting data
Municipality users can export aggregated data for all companies in the system by logging in and choosing a specific year to export from.

The data does not contain any company specific information and is intended for further analysis. The exported data is in `.csv` format using the Danish number formatting (comma as decimal separator) with semicolon as delimiter.

<img src="/docs/export.png" />
