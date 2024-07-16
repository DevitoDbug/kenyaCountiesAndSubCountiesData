# Kenyan Counties Data

This repository contains data for the counties in Kenya(i.e counties ,their respective headquarters and sub-counties), scraped using Golang and Colly.

## Data Source

The data includes the following fields for each county:
- Code: County code
- County: County name
- Headquarters: County headquarters
- SubCounties: Sub-counties within the county

The data was scraped from [Afro Blog - List of Counties](https://blog.afro.co.ke/list-of-counties/).

## Usage

You can download the JSON file containing the scraped data from the repository:
- [Download countiesDataFile.json](./countiesDataFile.json)

## About the Code

The scraping script (`main.go`) uses the Golang library Colly to fetch and parse HTML data from the website.

## Requirements

To run the script, you need:
- Golang installed on your machine
- Dependencies managed using Go modules (`go mod tidy`)

## Instructions

1. Clone the repository:
   ```bash
   git clone https://github.com/DevitoDbug/kenyaCountiesAndSubCountiesData.git
