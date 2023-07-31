# DatingApp
 - build with golang version 1.20
 - Echo framework
 - gorm
# How to Run
 - clone repo
 - go mod tidy
 - create new database with postgres, DB Name = DatingAppDB
 - go run main.go

# Dokumentation APi
- in folder doc_Api
# About App

1. User: Represents the user of the dating app. It includes attributes such as User ID, Username, Email, Password, and Premium Status.
2. Profile: Represents the dating profiles. It includes attributes such as Profile ID, User ID (foreign key to User), Name, Age, Gender, Location, and other relevant profile details.
3. Swipe: Represents the swipes made by users. It includes attributes such as Swipe ID, User ID (foreign key to User), Profile ID (foreign key to Profile), Swipe Type (like or pass), and Timestamp.

The relationships between these entities can be defined as:
- One User can have one or many Profiles (One-to-Many relationship).
- One Profile can be associated with only one User (One-to-One relationship).
- One User can make one or many Swipes (One-to-Many relationship).
- One Profile can be associated with many Swipes (One-to-Many relationship).

Sequence Diagram for the Dating Mobile App:
The Sequence Diagram illustrates the interaction and message flow between the objects in the system. Here's a textual representation of the sequence of events for the given functionalities:

1. User Sign up & Login:
- User sends a request to sign up with their credentials.
- The system verifies the uniqueness of the username and email.
- If the verification is successful, the user is registered in the system and receives a confirmation message.
- User logs in using their credentials.
- The system verifies the entered credentials and grants access to the user.

2. User Profile Viewing and Swiping:
- User requests to view dating profiles.
- The system retrieves 10 profiles for the user to view.
- User swipes left (pass) or swipes right (like) on a profile.
- The system records the swipe in the database, associating it with the user and the profile.
- The system checks if the profile has already been shown to the user on the same day.
- If the profile has been shown, the system retrieves a different profile for the user to view.
- This process continues until the user has viewed 10 profiles.

3. Premium Package Purchase:
- User selects a premium package to purchase.
- The system processes the payment transaction.
- The system updates the user's premium status and unlocks the chosen premium feature.