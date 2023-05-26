let post = {
    "Theme": null,
    "Title": null,
    "SubTitle": null,
    "AuthorName": null,
    "AuthorPhoto": null,
    "AuthorPhotoName": null,
    "Data": null,
    "BigImage": null,
    "BigImageName": null,
    "SmallImage": null,
    "SmallImageName": null,
    "Content": null
}

const
    readerAuthor = new FileReader();
    readerBig = new FileReader();
    readerSmall = new FileReader();

function Click()
{
    PrintToLog();
}    

function PrintToLog()
{
    let message = document.getElementById('message');
    if (message.childNodes.length !== 0 )
    {
        message.children[0].remove();
        message.children[0].remove();
    }
    let NotNull = true;
    console.log(post);

    for (key in post){
        if(post[key] === null)
        {
            NotNull = false;
        }
    }

    if (NotNull)
    {
        Preview();
        message.classList.remove('main-top__message_error');
        message.classList.add('main-top__message_complited');
        let icon = document.createElement('img');
        icon.src = "../static/sources/check_circle.svg";
        let text = document.createElement('text');
        text.textContent = "Publish Complete!"

        message.insertBefore(text, message.children[0]);
        message.insertBefore(icon, message.children[0]);
        message.children[0].classList.add('icon');

        let newpost = JSON.stringify(post);

        let XHR = new XMLHttpRequest();

        XHR.open("POST", "/api/post");
        XHR.send(newpost);
    }
    else
    {
        message.classList.remove('main-top__message_complited');
        message.classList.add('main-top__message_error');
        let icon = document.createElement('img');
        icon.src = "../static/sources/alert_circle.svg";
        let text = document.createElement('text');
        text.textContent = "Whoops! Some fields need your attention :o"

        
        
        message.insertBefore(text, message.children[0]);
        message.insertBefore(icon, message.children[0]);

        message.children[0].classList.add('icon');

    }
}

function Preview()
{
    let getTitleArticle = document.getElementById('preview-title-article');
        getTitle = document.getElementById('preview-title');
        getSubTitle = document.getElementById('preview-subtitle');
        getSubTitleArticle = document.getElementById('preview-subtitle-article');
        getAuthorName = document.getElementById('preview-author-name');
        getAuthorPhoto = document.getElementById('preview-author-photo');
        getData = document.getElementById('preview-data');
        getBigImage = document.getElementById('preview-big-image');
        getSmallImage = document.getElementById('preview-small-image');

    getTitleArticle.textContent = post.Title;
    getTitle.innerHTML = post.Title;
    getSubTitleArticle.innerHTML = post.SubTitle;
    getSubTitle.innerHTML = post.SubTitle;
    getAuthorName.innerHTML = post.AuthorName;
    getAuthorPhoto.src = post.AuthorPhoto;
    getData.innerHTML = post.Data;
    getBigImage.src = post.BigImage;
    getSmallImage.src = post.SmallImage;
}

function Select(name)
{
    let field = document.getElementById(name);
    field.classList.add('main-info__fields_select');
}

function NotSelect(name)
{
    let field = document.getElementById(name);
    let block = document.getElementById(name + '-block');
    field.classList.remove('main-info__fields_select');


    if((field.value === "") && (!field.classList.contains('main-info__fields_error')))
    {
        field.classList.add('main-info__fields_error');

        let err_message = document.createElement('p');
        if (name === 'Author-name')
        {
            err_message.textContent = 'Author name is required.';
        }
        else
        {
            err_message.textContent = name + ' is required.';
        }
        err_message.classList.add('main-info__error-message');

        block.insertBefore(err_message, block.children[2]);

        switch (name) {
            case 'Title':
                post.Title = null
                break;
            case 'Subtitle':
                post.SubTitle = null
                break;
            case 'Author-name':
                post.AuthorName = null
                break;
            case 'Theme':
                post.Theme = null
                break;
        }   
    }
    else
    {
        if(field.classList.contains('main-info__fields_error'))
        {
            field.classList.remove('main-info__fields_error');
            block.children[2].remove();
        }

        switch (name) {
            case 'Title':
                post.Title = field.value
                break;
            case 'Subtitle':
                post.SubTitle = field.value
                break;
            case 'Author-name':
                post.AuthorName = field.value
                break;
            case 'Theme':
                post.Theme = field.value
                break;    
        }    
    }
}


function ChangeIcon()
{
    let photo = document.getElementById('author-photo').files[0];
    let block = document.getElementById('author-photo-block');
    let label = document.getElementById('author-label');
    let removelabel = document.getElementById('remove-author');

    if ((post.AuthorPhoto !== "") && (removelabel.children[0] === undefined))
    {
        
        let icon = document.createElement('img');
        icon.id = "author-icon";
        icon.classList.add('main-info__photo-icon');

        block.insertBefore(icon, block.children[0]);

        label.children[0].remove();
        label.children[0].remove();
        
        let Upload = document.createElement('p');
        Upload.classList.add("main-info__upload");
        Upload.textContent = "Upload New";

        let remove = document.createElement('p');
        remove.classList.add("main-info__remove");
        remove.textContent = "Remove";

        let trash = document.createElement('img');
        trash.classList.add("main-info__icons");
        trash.src = "../static/sources/trash.svg";

        let camera = document.createElement('img');
        camera.classList.add("main-info__icons");
        camera.src = "../static/sources/camera.svg";

        label.insertBefore(Upload, label.children[0]);
        label.insertBefore(camera, label.children[0]);

        removelabel.insertBefore(remove, removelabel.children[0]);
        removelabel.insertBefore(trash, removelabel.children[0]);
    }

    let icon = document.getElementById('author-icon');

    if(photo)
    {
        readerAuthor.readAsDataURL(photo);
        post.AuthorPhotoName = photo.name;
    }

    readerAuthor.addEventListener(
        "load",
        ()=>{
            icon.src = readerAuthor.result;
            post.AuthorPhoto = readerAuthor.result;
        },
        false
    );
}

function RemoveIcon()
{
    let block = document.getElementById('author-photo-block');

    block.children[0].remove();

    let label = document.getElementById('author-label');
    label.children[0].remove();
    label.children[0].textContent = "Upload";

    let icon = document.createElement('img');
    icon.src = "../static/sources/Avatar.svg";
    icon.alt = "Avatar";

    label.insertBefore(icon, label.children[0]);

    let remove = document.getElementById('remove-author');
    remove.children[1].remove();
    remove.children[0].remove();

    post.AuthorPhoto = null;
}

function datachange()
{
    let data = document.getElementById('data');
    if(data.value === "")
    {
        post.Data = null;
    }
    else
    {
        post.Data = data.value
    }
}

function ChangeImage(size)
{
    let photo = document.getElementById(size + '-hero-image').files[0];
    let    label = document.getElementById(size + '-image-label');
    let    block = document.getElementById(size + '-block-image');
    let    removelabel = document.getElementById(size + '-remove-image');

    if((size === "big" ? post.BigImage !== "" : post.SmallImage !== "") && (removelabel.children[0] === undefined))
    {
        let image = document.createElement('img');
        image.id = size + '-image';

        block.insertBefore(image, block.children[0]);

        label.children[0].remove();

        let uploadnew = document.createElement('p');
        uploadnew.classList.add("main-info__upload");
        uploadnew.textContent = "Upload New";

        let remove = document.createElement('p');
        remove.classList.add("main-info__remove");
        remove.textContent = "Remove";

        let camera = document.createElement('img');
        camera.classList.add("main-info__icons");
        camera.src = "../static/sources/camera.svg";

        let trash = document.createElement('img');
        trash.classList.add("main-info__icons");
        trash.src = "../static/sources/trash.svg";

        

        label.insertBefore(uploadnew, label.children[0]);
        label.insertBefore(camera, label.children[0]);

        removelabel.insertBefore(remove, removelabel.children[0]);
        removelabel.insertBefore(trash, removelabel.children[0]);
    }

    let image = document.getElementById(size + '-image');

    if(size === "big")
    {
        if(photo)
        {
            readerBig.readAsDataURL(photo);
            post.BigImageName = photo.name;
        }
    
        readerBig.addEventListener(
            "load",
            ()=>{
                image.src = readerBig.result;
                post.BigImage = readerBig.result
            },
            false
        );

        image.classList.add('constructor__block__inp-style-icon-big')
    }
    else if(size === "small")
    {
        if(photo)
        {
            readerSmall.readAsDataURL(photo);
            post.SmallImageName = photo.name;

        }
    
        readerSmall.addEventListener(
            "load",
            ()=>{
                image.src = readerSmall.result;
                post.SmallImage = readerSmall.result
            },
            false
        );

        image.classList.add('main-info__small')
    }
}

function DeleteImage(size)
{
    let block = document.getElementById(size + '-block-image');
        label = document.getElementById(size + '-image-label');
        removelabel = document.getElementById(size + '-remove-image');
    
    block.children[0].remove();

    label.children[1].remove();
    label.children[0].remove();

    removelabel.children[1].remove();
    removelabel.children[0].remove();

    let image = document.createElement('img');

    image.src = "../static/sources/hero_image_" + size + ".png"
    image.alt = size + " image";

    label.insertBefore(image, label.children[0]);
}

function ChangeContent()
{
    let text = document.getElementById('content');

    if(text.value === "")
    {
        post.Content = null;
    }
    else
    {
        post.Content = text.value;
    }
}