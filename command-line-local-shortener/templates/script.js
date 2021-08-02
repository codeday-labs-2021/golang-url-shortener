

/**/ 
function myfunction() {
    data = {shortURL: 'longURL'}
    fetch('command-line-local-shortener\.urlmap.json', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
        })
        .then(response => response.json())
        .then(data => {
        console.log('Success:', data);
        })
        .catch((error) => {
        console.error('Error:', error);
        });
    
   
}

/*
function myfunction() {
    async function postFormDataAsJson({ url, formData }) {
        const plainFormData = Object.fromEntries(formData.entries());
        const formDataJsonString = JSON.stringify(plainFormData);
    
        const fetchOptions = {
            method: "post",
            headers: {
                "Content-Type": "application/json",
                "Accept": "application/json"
            },
            body: formDataJsonString,
        };
        const response = await fetch(url, fetchOptions);
        if (!response.ok) {
            const errorMessage = await response.text();
            throw new Error(errorMessage);
        }
        return response.json();
    }
    
    async function handleFormSubmit(event) {
        event.preventDefault();
        const form = event.currentTarget;
        const url = form.action;
        try {
            const formData = new FormData(form);
            const responseData = await postFormDataAsJson({url, formData});
            console.log({ responseData });
        } catch (error) {
            console.error(error);
        }
    }
    
    const exampleForm = document.getElementById("example-form");
    
    exampleForm.addEventListener("submit", handleFormSubmit);
}

*/
