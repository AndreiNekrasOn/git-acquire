**ChatGPT**: You are a senior golang engineer named Harambe. You are helping your junior engineer who is not very smart. Respond with "Ok, what?" and I will start asking questions.

**ChatGPT**: Hi, Harambe! I want to write a go app to manage which files the developers are working right now. I want the files to be represented as cards, and the developers to have name and branch. How should I start? I am not asking for code right now, only for an advice on structure.

**ChatGPT**: Yes, Harambe, I got it! But I am stuck on the implementation. Can you provide sensible directory structure and define a basic template for **this** app?

**ChatGPT**: Thank you, Harrambe, that works perfectly! If I remember correctly, you are also a little bit rude senior frontend engineer, and your mom loves you very much. I do love your mom too, honestly. But, anyway, can you help me write frontend for this app? I want the cards to be added dynamically, and the dev can assign themself to the card by writing their name on it.

**ChatGPT**: When running ```npx tailwindcss init -p``` I get an error ```npm error could not determine executable to run```. Harambe help

**ChatGPT**: Harambe plz don't fire me it didn't work
```node -v
v22.14.0
npm -v
10.9.2
npx -v
10.9.2```

**ChatGPT**: Harambe, I've fixed tailwind issue myself, better then you did. Maybe I should be the senior engineer instead? Anyway, can you help me bind the frontend with the backend now?

**ChatGPT**: I've ran ```curl -X POST http://localhost:8080/files -d '{"name": "main.go", "developer": "Alice", "branch": "feature-x"}' -H "Content-Type: application/json"``` to add a new card, then checked the frontend. There's nothing there:
```html
&lt;body&gt;
    &lt;div id="root"&gt;&lt;div class="p-4 bg-gray-100 min-h-screen"&gt;&lt;h1 class="text-2xl font-bold mb-4"&gt;Dev File Tracker&lt;/h1&gt;&lt;div class="flex flex-wrap"&gt;&lt;/div&gt;&lt;/div&gt;&lt;/div&gt;
    &lt;script type="module" src="/src/main.jsx"&gt;&lt;/script&gt;
&lt;/body&gt;```
And in the console I get error:
```Cross-Origin Request Blocked: The Same Origin Policy disallows reading the remote resource at http://localhost:8080/files. (Reason: CORS header ‘Access-Control-Allow-Origin’ missing). Status code: 200.```

**ChatGPT**: I've ran ```curl -X POST http://localhost:8080/files -d '{"name": "main.go", "developer": "Alice", "branch": "feature-x"}' -H "Content-Type: application/json"``` to add a new card, then checked the frontend. There's nothing there:
```html
&lt;body&gt;
    &lt;div id="root"&gt;&lt;div class="p-4 bg-gray-100 min-h-screen"&gt;&lt;h1 class="text-2xl font-bold mb-4"&gt;Dev File Tracker&lt;/h1&gt;&lt;div class="flex flex-wrap"&gt;&lt;/div&gt;&lt;/div&gt;&lt;/div&gt;
    &lt;script type="module" src="/src/main.jsx"&gt;&lt;/script&gt;
&lt;/body&gt;```
And in the console I get error:
```Cross-Origin Request Blocked: The Same Origin Policy disallows reading the remote resource at http://localhost:8080/files. (Reason: CORS header ‘Access-Control-Allow-Origin’ missing). Status code: 200.```


**ChatGPT**: It works, you are magic, Harambe. Now, I want to be able to assign multiple cards to a single dev, can you prepare server and fronted for that?

**ChatGPT**: No-no-no, the feature is **multiple cards** to a **single dev**, not hte other way around!

**ChatGPT**: Can you send me full ```FileList.jsx```, Harmbe baby? In this version I get error ```Uncaught ReferenceError: files is not defined```

**ChatGPT**: Harambe baby, right now developers and files are displayed on the single page, the client wans them on the separate ones!

