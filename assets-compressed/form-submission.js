const form = document.getElementById("orderForm");
const submitButton = form.querySelector('button[type="submit"]');

form.addEventListener("submit", async function(e) {
    e.preventDefault();
    
    // Disable submit button and show loading state
    submitButton.disabled = true;
    const originalText = submitButton.textContent;
    submitButton.textContent = "Submitting...";
    
    const formData = new FormData(form);
    
    try {
        const response = await fetch("/order", {
            method: "POST",
            body: formData
        });
        
        if (response.ok) {
            console.log("Form submitted successfully!");
            
            
            // Reset form
            form.reset();
            
        } else {
            console.error("Failed to submit form");
            alert("Failed to submit form. Please try again.");
        }
    } catch (err) {
        console.error("Error submitting form:", err);
        alert("Error submitting form. Please try again.");
    } finally {
        // Re-enable submit button
        submitButton.disabled = false;
        submitButton.textContent = originalText;
    }
});