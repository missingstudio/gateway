import { Header } from "../components/Header";
import Hero from "../components/Hero";
import styles from "./page.module.css";

export default function Page(): JSX.Element {
  return (
    <main className={styles.main}>
      <Header />
      <Hero />
    </main>
  );
}
