# sendx-backend-IIT2020140

# Web Crawler App

## **Table of Contents**
- [Overview](#overview)
- [Key Features](#key-features)
- [API Endpoints](#api-endpoints)
- [Installation & Setup](#installation--setup)
- [Demo](#demo)
  
## **Overview**

The Web Crawler App is an advanced, state-of-the-art web crawling solution engineered for efficiency and adaptability. Designed with scalability in mind, the application is equipped with a host of features to ensure seamless retrieval and delivery of web page content. One of its primary distinctions is its ability to prioritize paying customers, ensuring they receive optimal performance at all times.

## **Key Features**

### Prioritized Crawling
- **For Paying Customers:** Experience rapid crawling with minimal wait times.
- **For Non-Paying Customers:** Standard crawling with brief waiting periods ensures fairness without compromising quality.

### Configurable Concurrency
- **For Paying Customers:** Benefit from increased concurrency with up to 5 dedicated crawler workers.
- **For Non-Paying Customers:** Dedicated 2 crawler workers ensure reliable performance.

### Advanced Administrative Controls
Administrators possess comprehensive control:
- **Dynamic Worker Allocation:** Adjust the number of active crawler workers on-the-fly to match varying demands.
- **Customizable Crawling Speed:** Set and modify the rate at which pages are crawled, ensuring efficient use of resources.

### Intelligent Caching System
- **Optimized Performance:** Pages crawled within the past 60 minutes are stored in a cache to prevent redundant retrievals, accelerating response times.

## **API Endpoints**

- **Crawl Webpage:** `/crawl?url=YOUR_URL&isPaying=true/false` - Specify the target webpage and customer type for tailored performance.
- **Configure Worker Count:** `/config/numWorkers` - Dynamically set the count of active crawler workers via POST.
- **Set Crawling Rate:** `/config/maxCrawlsPerHour` - Define the crawl rate per worker every hour with POST.
- **Retrieve Configuration:** `/get-config` - Obtain the current server configurations in an organized JSON format.

## **Installation & Setup**

1. **Clone:** Initiate by cloning this repository to your local machine.
   ```bash
   git clone https://github.com/Ayushi181/sendx-backend-IIT2020140.git

## **Demo**

Witness the efficiency and versatility of the Web Crawler App in action:

- [**Watch the Project Demo Video**](https://drive.google.com/file/d/1LVUCYLg6XncXPeDPnQVRDXhAbWXk4-9P/view?usp=sharing).

