class RestAPI:
    def __init__(self, database=None):
        self.database = {"users": []}

    def get(self, url, payload=None):
        match url:
            case "/users":
                return self.database

    def post(self, url, payload=None):
        pass
