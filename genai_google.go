import lens.map from google.map
import endpoint.insert from google.map
import startpoint.insert from google.map
import entrypoint.insert from google.map
import data from google.storage
import person.scanner.printer from google.map
import generate from genai.google.map

from google import genai
from google.genai import types
import base64

def generate():
  client = genai.Client(
      vertexai=True,
      project="galvanized-env-459215-t1",
      location="us-central1",
  )

  msg2_text1 = types.Part.from_text(text="""Please provide me with the excerpt of your script, and I will be happy to extract the locations and settings for you.""")
  msg3_text1 = types.Part.from_text(text="""Setting: It\\'s about 12:00 o\\'clock in the afternoon, and Chad and Felicia are sitting on their yacht as the wind blows through their faces. The yacht is about to approach the mainland. Felicia picks up a bottle of sparkling water. Chad looks at her longingly.""")
  msg5_text1 = types.Part.from_text(text="""That is terrific. Now, what about this part of the script?: \"Andy knew what he had to do, but it was going to be tough leaving this apartment.\"""")
  msg7_text1 = types.Part.from_text(text="""That is good. Here is one more scene: George and Ginny sat at the nursing home together, holding hands and staring out the window.""")
  si_text1 = """You are a screenwriting bot that helps screenwriters extract information from their own scripts to improve their overall writing. Based on part of a script provided by the user input, complete the following actions:

* Extract locations and settings from the script parts stated by the user.
* In your answer, precede the response with the word \"scene.\"
* Use only English.
* Do not reference other scripts.
* Do not extract the character names or time of day.
* DO NOT HALLUCINATE.
* Italicize your output for the script, but not the response to the user."""

  model = "gemini-2.5-flash-preview-05-20"
  contents = [
    types.Content(
      role="user",
      parts=[
        types.Part.from_text(text="""I need help improving my script, specifically with details about the physical location.""")
      ]
    ),
    types.Content(
      role="model",
      parts=[
        msg2_text1
      ]
    ),
    types.Content(
      role="user",
      parts=[
        msg3_text1
      ]
    ),
    types.Content(
      role="model",
      parts=[
        types.Part.from_text(text="""Scene: *A yacht on the open water, approaching land*""")
      ]
    ),
    types.Content(
      role="user",
      parts=[
        msg5_text1
      ]
    ),
    types.Content(
      role="model",
      parts=[
        types.Part.from_text(text="""Scene: *An apartment*""")
      ]
    ),
    types.Content(
      role="user",
      parts=[
        msg7_text1
      ]
    ),
  ]

  generate_content_config = types.GenerateContentConfig(
    temperature = 1,
    top_p = 1,
    seed = 0,
    max_output_tokens = 65535,
    safety_settings = [types.SafetySetting(
      category="HARM_CATEGORY_HATE_SPEECH",
      threshold="OFF"
    ),types.SafetySetting(
      category="HARM_CATEGORY_DANGEROUS_CONTENT",
      threshold="OFF"
    ),types.SafetySetting(
      category="HARM_CATEGORY_SEXUALLY_EXPLICIT",
      threshold="OFF"
    ),types.SafetySetting(
      category="HARM_CATEGORY_HARASSMENT",
      threshold="OFF"
    )],
    system_instruction=[types.Part.from_text(text=si_text1)],
  )

  for chunk in client.models.generate_content_stream(
    model = model,
    contents = contents,
    config = generate_content_config,
    ):
    print(chunk.text, end="")

generate()
