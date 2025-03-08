import { FC } from 'react'
import { Link } from 'react-router-dom'
import styles from './ExperienceCard.module.css'

interface ExperienceCardProps {
  organization: string
  role: string
  duration: {
    start: string
    end: string
  }
  technologies: string[]
  location?: string
  logo?: string
}

const ExperienceCard: FC<ExperienceCardProps> = ({ 
  organization, 
  role, 
  duration,
  technologies,
  location,
  logo
}) => {
  // Function to get the correct image path
  const getImagePath = (path: string) => {
    if (path.startsWith('http')) return path;
    return import.meta.env.BASE_URL + path.replace(/^\//, '');
  };

  return (
    <Link to="/experience" className={styles.card}>
      <div className={styles.logo}>
        {logo ? (
          <img 
            src={getImagePath(logo)} 
            alt={`${organization} logo`} 
            className={styles.logoImage} 
          />
        ) : (
          <div className={styles.logoPlaceholder}>
            {organization.charAt(0)}
          </div>
        )}
      </div>
      
      <div className={styles.content}>
        <h3 className={styles.organization}>{organization}</h3>
        {location && <div className={styles.location}>{location}</div>}
        
        {/* ROLE VARIANT - Currently active */}
        <div className={styles.role}>{role}</div>

        {/* DURATION VARIANT - Uncomment this and comment the role variant above to switch*/}
        {/* {<div className={styles.duration}>{duration.start} - {duration.end}</div>} */}
      </div>
    </Link>
  )
}

export default ExperienceCard 