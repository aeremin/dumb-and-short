import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Dumb-and-Short",
  description: "Very simple and dumb URL shortener for private use",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        {children}
      </body>
    </html>
  );
}
