//data = {shortUrl:"", longUrl: longURL}
// obj = {shortUrl: "", longUrl: "http://google.com"}
var serializeForm = function (form) {
    var obj = {};
    var formData = new FormData(form);
    for (var key of formData.keys()) {
        obj[key] = formData.get(key);
    }
    console.log(obj);
    return obj;

};

function myfunction(event) {
    event.preventDefault();
    data = serializeForm(event.target); // turns form data into obj
    console.log(JSON.stringify(data));
    fetch('http://localhost:8080/create', {
        method: 'POST',
        body: JSON.stringify(data),
        headers: {
            'Accept': 'application/json',
            'Content-type': 'application/json'
        }
    }).then(function (response) {
        if (response.ok) {
            console.log(resonse.json);
            return response.json();
        }
        return Promise.reject(response);
    }).then(function (data) {
        console.log(data);
    }).catch(function (error) {
        console.warn(error);
        alert(error); //for debugging
    });
};



