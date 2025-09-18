import requests
import  json


def ask_model(prompt):
    url = "http://localhost:11434/api/generate"
    payload = {
        "model": "llama2",
    }
    response = requests.post(url, json=payload)
    for line in response.iter_lines():
        if line:
            print(json.loads(line)["response"], end="")
if __name__ ==" __main___":
    user_prompt = input("Enter your question for the model: ")
    ask_model(user_prompt)