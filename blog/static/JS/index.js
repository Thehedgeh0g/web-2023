function Click(){
    let header = document.getElementById('header')
    if (header.classList.contains('open')) {
        header.classList.remove('open');
    }
    else
    {
        header.classList.add('open');  
    }
}