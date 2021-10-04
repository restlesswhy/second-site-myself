let button = document.querySelector(".form-floating > button");
if (button) {
    button.onclick = function (e) {
        let inputs = document.querySelectorAll(".form-floating > input");

        let data = {};
        
        for (let i = 0; i < inputs.length; i++) {
            data[inputs[i].name] = inputs[i].value;
        }

        let xhr = new XMLHttpRequest();

        xhr.open("POST", "/user/auth");
        xhr.onload = function(e) {
            let response = JSON.parse(e.currentTarget.response);
            if ("Error" in response) {
                if (response.Error == null) {
                    console.log("User is authorizated!");
                    // window.location.url = "/";   
                } else {
                    console.log(response.Error);
                }
            } else {
                console.log("uncorrect auth");
            }
        };
        xhr.send(JSON.stringify(data));
    }
}   