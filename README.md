# agenda-maker
Agenda Maker for Mary

# Notes

The "Read for Meeting" section is generated based on the value of the column "Circulate". If the value for it on a given
row is "yes" then the row will be added to this section and removed from the preliminary report.

# Usage

- Download the latest release
- Open a new Terminal
- Navigate to the folder where you downloaded the release
- Place the CSV file that you want to use to generate the report in the same directory as the downloaded release.
- Run the binary by typing:
`./agenda-maker --file=NAMEOFTHEFILE.CSV`
Example:
`./agenda-maker --file=reports.csv`
