import requests

# Small example - paste in to Redash query

# Here you can use {{ param }}

# See https://www.energidataservice.dk/tso-electricity/ConsumptionDK3619codehour
response = requests.get(
    url='https://api.energidataservice.dk/dataset/ConsumptionDK3619codehour?limit=5')

result_json = response.json()
records = result_json.get('records', [])

# Define columns based on the fields you want to show
result = {
    "columns": [
        {"name": "HourDK", "type": "string", "friendly_name": "Hour (DK Time)"},
        {"name": "DK36Code", "type": "string", "friendly_name": "DK36 Code"},
        {"name": "DK36Title", "type": "number", "friendly_name": "DK36 Title"},
        {"name": "DK19Code", "type": "string", "friendly_name": "DK19 Code"},
        {"name": "DK19Title", "type": "string", "friendly_name": "DK19 Title"},
        {"name": "Consumption_MWh", "type": "number", "friendly_name": "Consumption (MWh)"},
        {"name": "Consumption_MWh_Comparison", "type": "string", "friendly_name": "Consumption MWh Comparison"}
    ],
    "rows": []
}

# Populate rows
for record in records:
    result["rows"].append({
        "HourDK": str(record.get("HourDK")),
        "DK36Code": str(record.get("DK36Code")),
        "DK36Title": record.get("DK36Title"),
        "DK19Code": str(record.get("DK19Code")),
        "DK19Title": str(record.get("DK19Title")),
        "Consumption_MWh": str(round(float(record.get("Consumption_MWh")) / 5)),
        "Consumption_MWh_Comparison": "4 over gennemsnit for din branche",
    })