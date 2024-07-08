// Updated JavaScript code
document.addEventListener('DOMContentLoaded', function() {
    console.log('JavaScript loaded.');

    // Example: Changing background color on click
    const body = document.querySelector('body');
    body.addEventListener('click', function() {
        body.style.backgroundColor = '#f39c12';
    });

    // Example: Displaying an alert on hover
    const h1 = document.querySelector('h1');
    h1.addEventListener('mouseover', function() {
        alert('Hovered over the heading!');
    });
});
