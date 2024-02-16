import ProvidersList from "./providers-list";

export const metadata = {
  title: "Providers - AI studio",
  description: "List of providers",
};

export default function Providers() {
  return (
    <>
      <div className="container flex flex-col gap-4 p-4">
        <ProvidersList />
      </div>
    </>
  );
}
