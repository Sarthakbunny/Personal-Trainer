# 💪 Personalized Diet & Workout Suggestion App – Project Overview

---

## 🧭 Objective

Build a lightweight web application that delivers **personalized workout and diet recommendations** by asking a few smart questions during onboarding. Based on the responses, the app uses **AI (Gemini)** to generate daily actionable suggestions.

The aim is to deliver a functional MVP **within one week**, focusing on user personalization, fast responses, and clean UI.

---

## 🎯 Core Features

* ✅ **Onboarding Questionnaire**
  Collect information like fitness goals, dietary preferences, workout habits, etc.

* 🧠 **AI-Powered Recommendations**
  Use **Gemini API** to generate personalized daily workout and diet plans.

* 📅 **Daily Tracker / Streak Counter**
  Let users track whether they followed the day’s plan (check-in system).
  Encourages consistency and shows streak stats on dashboard.

* 🏠 **User Dashboard**
  Display:

  * Today’s workout and diet
  * Option to regenerate
  * Streak stats

* 🔐 **Authentication**
  Firebase Auth (anonymous or email login)

---

## 🧰 Tech Stack

| Layer          | Tech                                               | Notes                                                   |
| -------------- | -------------------------------------------------- | ------------------------------------------------------- |
| Backend        | **Go (Golang)**                                    | Fast, scalable, easy concurrency                        |
| Frontend       | **React.js**                                       | Modern UI, reusable components                          |
| Database       | **Firebase Firestore**                             | Real-time DB, easy setup, integrates with Firebase Auth |
| AI Integration | **Gemini API**                                     | Context-aware suggestions using natural language        |
| Hosting        | **Vercel (Frontend)**, **Render/Fly.io (Backend)** | Easy, quick deployment                                  |

---

## 📅 7-Day Development Plan

| Day       | Task                                                                   |
| --------- | ---------------------------------------------------------------------- |
| **Day 1** | Set up project structure (frontend + backend), design basic wireframes |
| **Day 2** | Build onboarding questionnaire UI, setup Firebase Auth                 |
| **Day 3** | Create Go backend with API routes, connect to Firestore                |
| **Day 4** | Integrate Gemini API for generating suggestions                        |
| **Day 5** | Build dashboard: display daily plan + streak tracking                  |
| **Day 6** | Add check-in system, polish UI, handle errors                          |
| **Day 7** | Final testing, deployment, prepare readme/docs                         |

---

## 🚀 Post-MVP Ideas

* Save & view suggestion history
* Push notifications/reminders for check-in
* Advanced streak analytics (graphs, longest streak, etc.)
* Optional community leaderboard
* Add other AI models (OpenAI, Claude)

---

## 🔁 Alternatives (Optional)

| Use Case        | Option         | Why                                    |
| --------------- | -------------- | -------------------------------------- |
| DB              | Supabase       | If you want SQL + row-level security   |
| Auth            | Clerk/Auth0    | For better UI & user management        |
| AI              | OpenAI, Claude | More control over tone, output quality |
| Backend Hosting | Railway        | Easiest Go deployment with GitHub CI   |
