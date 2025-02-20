        // Listen for the htmx:wsAfterMessage event
        document.body.addEventListener('htmx:wsAfterMessage', function(event) {
            // Get the message content from the event detail
            const messageContent = event.detail.message;

            // Get the template
            const template = document.getElementById('message-template');

            // Clone the template content
            const messageElement = template.content.cloneNode(true);
            const myDiv = messageElement.querySelector("p")
            // Replace the placeholder with the actual message
            myDiv.textContent = messageContent;

            // Append the new message element to the container
            document.getElementById('message-container').appendChild(messageElement);
        });

