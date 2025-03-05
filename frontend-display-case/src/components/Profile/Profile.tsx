import { FC } from 'react'
import styles from './Profile.module.css'

const Profile: FC = () => {
  return (
    <section className={styles.profile}>
    <div className={styles.content}>
      <div className={styles.textContent}>
        <h1 className={styles.name}>Your Name</h1>
        <h2 className={styles.title}>Software Engineer</h2>
        <div className={styles.links}>
          <a href="https://github.com/yourusername" target="_blank" rel="noopener noreferrer">
            GitHub
          </a>
          <span className={styles.separator}>|</span>
          <a href="https://linkedin.com/in/yourusername" target="_blank" rel="noopener noreferrer">
            LinkedIn
          </a>
          <span className={styles.separator}>|</span>
          <a href="mailto:your.email@example.com">
            Email
          </a>
        </div>
        <div className={styles.summary}>
          <p>
            A passionate software engineer with X years of experience in building scalable applications.
            Specialized in full-stack development with expertise in React, TypeScript, and Go.
            Currently focused on cloud-native solutions and distributed systems.
          </p>
        </div>
      </div>
      <div className={styles.imageContainer}>
        <img 
          src="/headshot.jpg" 
          alt="Professional headshot" 
          className={styles.profileImage}
        />
      </div>
    </div>
  </section>
  )
}

export default Profile