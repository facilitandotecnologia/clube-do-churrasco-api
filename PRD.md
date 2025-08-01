# Product Requirements Document (PRD)  
**Product:** Clube do Churrasco  
**Purpose:** Organize barbecues with friends

---

## 1. Overview

Clube do Churrasco is a mobile/web application designed to help users organize barbecue events with friends. The app streamlines event creation, guest management, payment requests and confirmations, user account management, and helps users consult meat prices at local meat houses.

---

## 2. Features

### 2.1. User Account Management (CRUD)
- Users can **Create** an account with required information (name, email, WhatsApp number, password).
- Users can **Read/View** their account details and profile.
- Users can **Update** their account information (e.g., change name, email, WhatsApp number, password).
- Users can **Delete** their account, removing all personal data and event associations.

### 2.2. Create Event
- Users can create a new barbecue event.
- Required fields:  
  - **Date:** When the event will take place.  
  - **Description:** Brief summary of the event.  
  - **Attenders:** List of friends invited.  
  - **Event Details:** Additional notes or description.

### 2.3. Add Friends to Event
- Users can search for and add friends to the event as attenders.
- Friends receive notifications/invitations.

### 2.4. Consult Meat Prices
- Users can browse or search for meat prices at local meat houses.
- Integration with external APIs or manual price entry by meat houses.

### 2.5. Request Payment from Attenders
- The event admin can request payment from each attender for event costs.
- Payment requests are sent individually to each attender.
- Attenders receive payment requests in-app and/or via WhatsApp message.
- Payment status is tracked per attender.

### 2.6. Attender Sends Payment Proof & Real-Time Recognition
- Attenders can upload/send payment proof (e.g., payment receipt or bill) to the admin via the app.
- The app processes and recognizes the payment in real time (e.g., via integration with payment gateways or automated receipt analysis).
- Admin is notified instantly when a payment is recognized and confirmed.

---

## 3. User Stories

- **As a user,** I want to create, view, update, and delete my account so I can manage my personal information and privacy.
- **As a user,** I want to create a barbecue event so I can organize gatherings with my friends.
- **As a user,** I want to add friends to my event so they can be invited and participate.
- **As a user,** I want to check meat prices at local stores so I can plan my shopping and budget.
- **As an event admin,** I want to request payment from attenders so that event costs are shared.
- **As an attender,** I want to receive payment requests in the app or via WhatsApp so I know how much to pay and how.
- **As an attender,** I want to send my payment proof to the admin via the app so my payment can be recognized and confirmed in real time.
- **As an admin,** I want to be notified instantly when an attender's payment is recognized.

---

## 4. Data Model

- **User**
  - id
  - name
  - email
  - WhatsApp number
  - password (hashed)
- **Event**
  - id
  - date
  - description
  - details
  - attenders (list of user IDs)
  - admin (user ID)
  - paymentRequests (list of payment request objects)
- **MeatHouse**
  - id
  - name
  - location
  - meatPrices (list of meat types and prices)
- **PaymentRequest**
  - id
  - eventId
  - attenderId
  - amount
  - status (pending, paid, overdue)
  - sentVia (in-app, WhatsApp)
  - paymentProof (file/image or transaction ID)
  - recognized (boolean)
  - recognizedAt (timestamp)

---

## 5. Non-Functional Requirements

- Responsive design for mobile and web.
- Secure authentication for users.
- Real-time updates for invitations, payment requests, payment recognition, and price changes.
- Integration with WhatsApp API for sending payment requests.
- Integration with payment gateways or receipt analysis for real-time payment recognition.
- GDPR-compliant data deletion for user accounts.

---

## 6. Out of Scope

- Payment processing (actual transfer of funds).
- Delivery or logistics management.

---