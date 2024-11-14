export default function Auth({
    children,
  }: Readonly<{
    children: React.ReactNode;
  }>) {
    return (
      <main>
        {children}
      </main>
    );
  }