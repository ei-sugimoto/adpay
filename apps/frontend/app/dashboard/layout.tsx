import PageGuard from '@/utils/pageGuard';

export default function DashboardLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return <PageGuard>{children}</PageGuard>;
}
