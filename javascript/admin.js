async function updateConfig() {
    const numWorkers = document.getElementById('numWorkers').value;
    const maxCrawls = document.getElementById('maxCrawls').value;

    // Check if the input fields are not empty and are valid numbers
    if (!numWorkers && !maxCrawls && isNaN(numWorkers) && isNaN(maxCrawls)) {
        alert("Please enter valid numbers in both fields.");
        return;
    }

    await fetch('/config/numWorkers', {
        method: 'POST',
        body: numWorkers,
        headers: { 'Content-Type': 'application/json' }
    });

    await fetch('/config/maxCrawlsPerHour', {
        method: 'POST',
        body: maxCrawls,
        headers: { 'Content-Type': 'application/json' }
    });

    // Show alert with configuration values
    alert(`Configuration updated successfully!\n\nNumber of Workers: ${numWorkers}\nMax Crawls per Hour: ${maxCrawls}`);

    // Clear the input fields
    document.getElementById('numWorkers').value = '';
    document.getElementById('maxCrawls').value = '';
}

async function getConfig() {
    const response = await fetch('/config');
    const data = await response.json();
    document.getElementById('configOutput').textContent = JSON.stringify(data, null, 2);
}

function clearConfig() {
    document.getElementById('numWorkers').value = '';
    document.getElementById('maxCrawls').value = '';
}

async function getCurrentConfig() {
    try {
        // Fetch data from the /get-config endpoint
        const response = await fetch('/get-config');

        // If the fetch request is not successful, throw an error
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        // Parse the JSON response
        const configData = await response.json();

        // Update the input fields with the configuration data
        document.getElementById('numWorkers').value = configData.numWorkers;
        document.getElementById('maxCrawls').value = configData.maxCrawlsPerHour;
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error.message);
    }
}


function escapeHtml(unsafe) {
    return unsafe
        .replace(/&/g, "&amp;")
        .replace(/</g, "&lt;")
        .replace(/>/g, "&gt;")
        .replace(/"/g, "&quot;")
        .replace(/'/g, "&#039;");
}