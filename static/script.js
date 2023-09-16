// JavaScript code for submitting the form via AJAX
document.getElementById("bookForm").addEventListener("submit", function (event) {
    event.preventDefault();
    const formData = new FormData(this);

    fetch("/create-book", {
        method: "POST",
        body: formData,
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert("Error: " + data.error);
        } else {
            alert("Book created successfully!");
            location.reload();
        }
    })
    .catch(error => {
        console.error("Error:", error);
        alert("An error occurred while creating the book.");
    });
});
