function getQueryParam(param) {
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    return urlParams.get(param);
}

document.addEventListener("DOMContentLoaded", function() {
    const errorValue = getQueryParam('error');
    const checkValue = getQueryParam('check');
    const passwordDiv = document.getElementById('password');

    if (errorValue === 'badpassword') {
        console.log('Mot de passe incorrect');
        passwordDiv.textContent = 'Mot de passe incorrect';
        passwordDiv.style.color = 'red';
        passwordDiv.style.fontSize = '10px'; 
    } 

    if (checkValue === 'passwordchange') {
        console.log('Mot de passe correct');
        passwordDiv.textContent = 'Mot de passe bien changé';
        passwordDiv.style.color = 'green';
        passwordDiv.style.fontSize = '10px'; 
    } 
});
