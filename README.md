# Voicy
A little Slack integration for sending short voice messages via Slack


Currently there's not much to see here, except for a quick hack-together sample app that provides a Web interface to record audio as MP3 and a backend service (in Golang) to serve the assets and receive the audio.

In the end, **voicy** should be able to do the following:

* Provide a popup or interactive Slack channel area to record and submit voice-audio (only visible to sender)
* Provide a clickable Slack channel area to listen to the recorded audio (visible to all in the channel)
* Provide a transcript of the voice recording so that users can also read up on the conversation (after all, Slack is a text-chat kind of app)

* This should be usable in the XING context which means, that the transcription service most-probably cannot be a public one (Google/AWS text-to-speech)
* It should be runnable within our MISC kubernetes cluster