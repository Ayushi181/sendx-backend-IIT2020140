function login() {
    const userType = document.querySelector('input[name="user-type"]:checked').value;
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    if (userType === 'user') {
        window.location.href = `/user-login`;
        // alert("User Login Successful");
    } else if (userType === 'admin') {
        if (username === 'ayushi' && password === 'ayushi') {
            window.location.href = `/admin-login`;
            // alert("Admin Login Successful");
        } else {
            alert("Incorrect username or password");
        }
    }
}