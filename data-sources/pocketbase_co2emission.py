import requests

token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2xsZWN0aW9uSWQiOiJwYmNfMzE0MjYzNTgyMyIsImV4cCI6MTg1OTgyMzA4MywiaWQiOiJhMXpzMmQyMXg2ajhhb2oiLCJyZWZyZXNoYWJsZSI6ZmFsc2UsInR5cGUiOiJhdXRoIn0.KLPE9rc6Sy8i6x38eSO8Xnnm4rA6inQ83wi4URUSvaM"

response = requests.get(
        url='http://dev:8080/api/collections/co2_emission/records?perPage=9999',
    headers={
        "Authorization": f"Bearer {token}",
        "Content-Type": "application/json",
    }
)

result_json = response.json()
records = result_json.get('items', [])


# Define columns based on the fields you want to show
result = {
    "columns": [
        {"name": "value", "type": "float", "friendly_name": "value"},
        {"name": "organization", "type": "string", "friendly_name": "Organization"},
        {"name": "year", "type": "string", "friendly_name": "Year"},
        {"name": "domain", "type": "string", "friendly_name": "Domain"},
        {"name": "unit", "type": "string", "friendly_name": "Unit"},
        {"name": "type", "type": "string", "friendly_name": "type"},
        {"name": "aggregation", "type": "string", "friendly_name": "Aggregation"}
    ],
    "rows": []
}

# Populate rows
for record in records:
    result["rows"].append({
        "value": float(record.get("value")),
        "organization": str(record.get("organization")),
        "year": str(record.get("year")),
        "domain": str(record.get("domainName")),
        "unit": str(record.get("unitName")),
        "type": str(record.get("typeName")),
        "aggregation": str(record.get("timeAggregation"))
    })

