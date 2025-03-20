import { FC } from 'react'
import styles from './Footer.module.css'
import content from '../../data/content.json'

const Footer: FC = () => {
  const currentYear = new Date().getFullYear()
  const { profile } = content

  return (
    <footer className={styles.footer}>
      <p className={styles.copyright}>
        © {currentYear} made with a lot of ☕ by {profile.name}. All rights reserved.
      </p>
    </footer>
  )
}

export default Footer 