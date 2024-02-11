import Image from "next/image";
import Link from "next/link";
import LogoImg from "public/images/logo.svg";

export default function Logo() {
  return (
    <Link className="block" href="/" aria-label="Missing studio">
      <Image src={LogoImg} height={24} priority alt="Missing studio" />
    </Link>
  );
}
