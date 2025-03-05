import { FC } from 'react'
import styles from './Profile.module.css'

interface ProfileLink {
  url: string
  label: string
}

interface ProfileData {
  name: string
  title: string
  links: {
    github: ProfileLink
    linkedin: ProfileLink
    email: ProfileLink
  }
  summary: string
  image: string
}

interface ProfileProps {
  data: ProfileData
}

const Profile: FC<ProfileProps> = ({ data }) => {
  return (
    <section className={styles.profile}>
      <div className={styles.content}>
        <div className={styles.textContent}>
          <h1 className={styles.name}>{data.name}</h1>
          <h2 className={styles.title}>{data.title}</h2>
          <div className={styles.links}>
            <a href={data.links.github.url} target="_blank" rel="noopener noreferrer">
              {data.links.github.label}
            </a>
            <span className={styles.separator}>|</span>
            <a href={data.links.linkedin.url} target="_blank" rel="noopener noreferrer">
              {data.links.linkedin.label}
            </a>
            <span className={styles.separator}>|</span>
            <a href={data.links.email.url}>
              {data.links.email.label}
            </a>
          </div>
          <div className={styles.summary}>
            <p>{data.summary}</p>
          </div>
        </div>
        <div className={styles.imageContainer}>
          <img 
            src={data.image} 
            alt="Professional headshot" 
            className={styles.profileImage}
          />
        </div>
      </div>
    </section>
  )
}

export default Profile