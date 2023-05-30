let user = {
    "Email": null,
    "Password": null
}

let XHR = new XMLHttpRequest();

function password_visibility(){
    password = document.getElementById('password-field');

    if(password.type === 'text')
    {
        password.type = 'password';
    }
    else
    {
        password.type = 'text';
    }
}

function Select(type)
{
    let field = document.getElementById(type + '-field');

    if(!field.classList.contains('login-box__field_error'))
    {
        field.classList.add('login-box__field_select');   
    }
}

function NotSelect(type)
{
    let field = document.getElementById(type + "-field");
    let block = document.getElementById(type + "-block");
    
    field.classList.remove('login-box__field_select')

    if((field.value === "") && (!field.classList.contains('login-box__field_error')))
    {
        field.classList.add('login-box__field_error');

        let label = document.createElement('p');

        switch (type) {
            case 'login':
                label.textContent = "Email is required.";
                break;
            case 'password':
                label.textContent = "Password is required.";
                break;
        }

        label.classList.add('login-box__error');

        block.insertBefore(label, block.children[2]);
    }
    else if(!(field.value === ""))
    {
        if(field.classList.contains('login-box__field_error'))
        {
            field.classList.remove('login-box__field_error');
            block.children[2].remove();
        }

        switch (type) {
            case 'login':
                user.Email = field.value;
                break;
            case 'password':
                user.Password = field.value;
                break;
        }
    }
}
function DataError()
{
    let message = document.getElementById('message');
    if(!message.classList.contains('login-box__message'))
    {
        let icon = document.createElement('img');
        icon.classList.add('login-box__icon');
        icon.src = "../static/sources/alert_circle.svg";
        
        let text = document.createElement('p');
        text.classList.add('login-box__message-text');
        text.textContent = "A-Ah! Check all fields";
    
        message.classList.add('login-box__message');
        message.insertBefore(text, message.children[0]);
        message.insertBefore(icon, message.children[0]);
    }
}

function Click()
{
    if((user.Email === null) && (user.Password === null))
    {
        DataError();
        NotSelect('login');
        NotSelect('pass');  
    }
    else
    {
        let email = user.Email;
        let email_name = email.slice(0, email.indexOf('@')); 
        let is_email_valid = false;
    
        if((email.includes('@') && (email.indexOf('@') == email.lastIndexOf('@'))) && 
        (email.includes('.') && (email.lastIndexOf('.') - email.indexOf('@') > 1)))
        {
            is_email_valid = true;
        }
        
        if((!(email_name !== "")) || (!is_email_valid))
        {
            is_email_valid = false;
    
            let field = document.getElementById('login-field');
            if(!field.classList.contains("login-box__field_error"))
            {
                field.classList.add('login-box__field_error');
                
                let err_message = document.createElement('p');
                err_message.classList.add('login-box__error');
                err_message.textContent = "Incorrect email format. Correct format is ****@**.***";
        
                let block = document.getElementById('login-block');
                block.insertBefore(err_message, block.children[2]);
            }

            DataError();
        }
            
        if(is_email_valid)
        {
            let userdata = JSON.stringify(user);

            XHR.open("POST", "/api/login");
            XHR.send(userdata);
        }
    }
}

XHR.onload = function() {
    if (XHR.status != 200){
        DataError();
    }
    else {
        window.location.href = "/admin"
    }
}