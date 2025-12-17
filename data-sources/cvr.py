import requests

# Example code, needs to be fetched from the correct source

# cvr_number = "24213366"  # Example CVR number

# url = f"https://cvrapi.dk/api" or f"https://datacvr.virk.dk/gateway/virksomhed/hentVirksomhed?cvrnummer=24213366&locale=da"
# params = {
#     "search": cvr_number,
#     "country": "dk"
# }

url = f"https://datacvr.virk.dk/gateway/virksomhed/hentVirksomhed?cvrnummer=24213366&locale=da"

# response = requests.get(url, params=params)
# response = requests.get(url)
# response.raise_for_status()
# data = response.json()

data = {
    'vat': 24213366,
    'name': 'ALEXANDRA INSTITUTTET A/S',
    'address': 'Åbogade 34',
    'zipcode': '8200',
    'city': 'Aarhus N',
    'cityname': None,
    'protected': True,
    'phone': None,
    'email': None,
    'fax': None,
    'startdate': '27/08 - 1999',
    'enddate': None,
    'employees': 103,
    'addressco': None,
    'industrycode': 721000,
    'industrydesc': 'Forskning og eksperimentel udvikling inden for naturvidenskab og teknik',
    'companycode': 60,
    'companydesc': 'Aktieselskab',
    'creditstartdate': None,
    'creditbankrupt': False,
    'creditstatus': None,
    'owners': None,
    'productionunits': [
        {
            'pno': 1006413582,
            'main': True,
            'name': 'ALEXANDRA INSTITUTTET A/S',
            'address': 'Åbogade 34',
            'zipcode': '8200',
            'city': 'Aarhus N',
            'cityname': None,
            'protected': True,
            'phone': None,
            'email': None,
            'fax': None,
            'startdate': '27/08 - 1999',
            'enddate': None,
            'employees': 73,
            'addressco': None,
            'industrycode': 702000,
            'industrydesc': 'Virksomhedsrådgivning og anden ledelsesrådgivning'
        },
        {
            'pno': 1012676375,
            'main': False,
            'name': 'ALEXANDRA INSTITUTTET A/S',
            'address': 'Rued Langgaards Vej 7, 5',
            'zipcode': '2300',
            'city': 'København S',
            'cityname': None,
            'protected': True,
            'phone': '70277091',
            'email': 'alexandra@alexandra.dk',
            'fax': None,
            'startdate': '01/01 - 2005',
            'enddate': None,
            'employees': 30,
            'addressco': None,
            'industrycode': 721000,
            'industrydesc': 'Forskning og eksperimentel udvikling inden for naturvidenskab og teknik'
        }
    ],
    't': 100,
    'version': 6
}

result = {
  "columns": [
    {"name": "cvr_number", "type": "string"},
    {"name": "company_name", "type": "string"},
    {"name": "address", "type": "string"},
    {"name": "postal_code", "type": "string"},
    {"name": "city", "type": "string"},
    {"name": "branche", "type": "string"},
  ],
  "rows": [
    {
      "cvr_number": str(data.get("vat")),
      "company_name": data.get("name"),
      "address": data.get("productionunits")[0].get("address"),
      "postal_code": data.get("zipcode"),
      "city": data.get("city"),
      "branche": str(data.get("industrycode")),
    }
  ]
}