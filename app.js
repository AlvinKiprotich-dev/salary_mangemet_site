// Function to render the Dashboard
function renderDashboard(username) {
    const contentDiv = document.getElementById('content');
    contentDiv.innerHTML = `
        <h2>Welcome, ${username}!</h2>
        <div class="salary-summary">
            <h3>Income Summary</h3>
            <p>Daily Income: $<span id="daily-income">0</span></p>
            <p>Fortnight Income: $<span id="fortnight-income">0</span></p>
            <p>Monthly Income: $<span id="monthly-income">0</span></p>
        </div>
        <div class="work-log">
            <h3>Work Logs</h3>
            <table>
                <thead>
                    <tr>
                        <th>Date</th>
                        <th>Hours Worked</th>
                        <th>Overtime</th>
                        <th>Daily Income</th>
                    </tr>
                </thead>
                <tbody id="worklog-table">
                    <!-- Rows will be dynamically added here -->
                </tbody>
            </table>
        </div>
    `;

    // TODO: Fetch data from backend and populate Dashboard
    populateDashboard();
}

// Replace the Login Form submission handler
loginForm.addEventListener('submit', (e) => {
    e.preventDefault();
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    // TODO: Authenticate user with the backend
    if (username && password) {
        renderDashboard(username);
    } else {
        alert('Please enter valid credentials.');
    }
});

// Function to populate dashboard data
function populateDashboard() {
    // Mock data for now
    const workLogs = [
        { date: '2024-11-01', hours: 8, overtime: 2, income: 120 },
        { date: '2024-11-02', hours: 7, overtime: 1, income: 105 },
    ];

    let dailyTotal = 0;
    let fortnightTotal = 0;
    let monthlyTotal = 0;

    const worklogTable = document.getElementById('worklog-table');
    workLogs.forEach((log) => {
        dailyTotal += log.income;
        fortnightTotal += log.income;
        monthlyTotal += log.income;

        worklogTable.innerHTML += `
            <tr>
                <td>${log.date}</td>
                <td>${log.hours}</td>
                <td>${log.overtime}</td>
                <td>$${log.income}</td>
            </tr>
        `;
    });

    // Update income summaries
    document.getElementById('daily-income').textContent = dailyTotal.toFixed(2);
    document.getElementById('fortnight-income').textContent = fortnightTotal.toFixed(2);
    document.getElementById('monthly-income').textContent = monthlyTotal.toFixed(2);
}
