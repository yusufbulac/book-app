# Case Study - Frontend

This is the **frontend** implementation for the ByFood case study, built with **Next.js 15 (App Router)**, **TypeScript**, and **Tailwind CSS**.

## Tech Stack
- **Next.js 15** (App Router)
- **TypeScript**
- **Tailwind CSS**
- **ESLint** + **TypeScript ESLint**
- **Context API** for global state management
- **React Hook Form** for form handling
- **Zod** for form validation
- **Toast notifications**
- **Responsive UI** similar to ByFood design style

## Features
- **Book Listing**
- **Book Creation** via modal form
- **Book Update** via modal form
- **Book Deletion** with confirmation modal
- **Book Details Page**
- **Sort & Filter** books (managed via Context API)
- **Toast Notification System**
- **Custom 404 Page** (`not-found.tsx`)
- **Loading State** (`loading.tsx`)
- **Error Page** (`error.tsx`)
- **Fully responsive layout**

## Project Structure
```
src/
 ├── app/                # App Router pages
 │    ├── books/          # Book detail pages
 │    ├── error.tsx       # Error page
 │    ├── loading.tsx     # Loading state
 │    ├── not-found.tsx   # 404 page
 │    ├── globals.css     # Global styles
 │    └── layout.tsx      # Root layout
 ├── components/         # Reusable UI components
 ├── context/            # Context API providers
 ├── lib/                # Helper utilities
 ├── services/           # API service functions
 ├── types/              # TypeScript types
 └── styles/             # Tailwind config & global CSS
```

## Installation & Setup
```bash
# Install dependencies
npm install

# Run development server
npm run dev

# Lint and auto-fix issues
npm run lint:fix

# Type check
npm run typecheck
```

## Environment Variables
Create a `.env.local` file in the `frontend/` directory with:
```
NEXT_PUBLIC_API_BASE_URL=http://localhost:8080
```

## Scripts
- `npm run dev` – Start dev server
- `npm run build` – Build for production
- `npm run start` – Start production server
- `npm run lint` – Run ESLint
- `npm run lint:fix` – Run ESLint with autofix
- `npm run typecheck` – Run TypeScript type check

## Notes
- This project follows **ESLint** and **TypeScript** best practices.
- The design is inspired by ByFood's layout and styling.
- Context API is used for state management of sorting/filtering across components.

---
**Author:** Yusuf Bulaç
