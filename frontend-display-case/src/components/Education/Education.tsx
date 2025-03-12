import { FC, useState } from 'react'
import ReactMarkdown from 'react-markdown'
import styles from './Education.module.css'

interface EducationItem {
  institution: string
  description: string
  location?: string
  duration: {
    start: string
    end: string
  }
  degree: string
  highlights?: string[] // Academic highlights instead of achievements
  technologies?: string[] // Optional for technologies/tools learned
}

interface EducationProps {
  title: string 
  items: EducationItem[]
}

const Education: FC<EducationProps> = ({ title, items }) => {
  const [expandedHighlights, setExpandedHighlights] = useState<Record<number, boolean>>({});

  const toggleHighlights = (itemIndex: number) => {
    setExpandedHighlights(prev => ({
      ...prev,
      [itemIndex]: !prev[itemIndex]
    }));
  };

  return (
    <section className={styles.education}>
      <h2 className={styles.sectionTitle}>{title}</h2>
      <div className={styles.timeline}>
        {items.map((item, index) => (
          <div key={index} className={styles.timelineItem}>
            <div className={styles.header}>
              <h3 className={styles.institution}>
                {item.institution} <span className={styles.separator}>|</span> <span className={styles.degree}>{item.degree}</span>
              </h3>
            </div>
            
            <div className={styles.duration}>
              {item.duration.start} - {item.duration.end}
            </div>

            {item.location && (
              <div className={styles.location}>{item.location}</div>
            )}

            <p className={styles.description}>{item.description}</p>

            {item.technologies && (
              <div className={styles.technologies}>
                <span>Technologies & Tools: </span>
                {item.technologies.map((tech, i) => (
                  <span key={i} className={styles.tech}>
                    {tech}
                  </span>
                ))}
              </div>
            )}

            {/* Render academic highlights */}
            {item.highlights && item.highlights.length > 0 && (
              <div className={styles.highlightsList}>
                <div className={styles.highlightContainer}>
                  <button
                    className={styles.highlightToggle}
                    onClick={() => toggleHighlights(index)}
                    aria-expanded={expandedHighlights[index]}
                    aria-controls={`highlights-${index}`}
                  >
                    <div className={styles.highlightName}>
                      Academic Highlights
                      <span className={styles.toggleIcon} aria-hidden="true">
                        {expandedHighlights[index] ? '−' : '+'}
                      </span>
                    </div>
                  </button>
                  
                  <div 
                    id={`highlights-${index}`}
                    className={`${styles.expandableContent} ${expandedHighlights[index] ? styles.expanded : ''}`}
                  >
                    <ul className={styles.highlights}>
                      {item.highlights.map((highlight, i) => (
                        <li key={i}>
                          <ReactMarkdown
                            components={{
                              strong: ({node, ...props}) => <strong className={styles.markdown} {...props} />,
                              a: ({node, ...props}) => <a className={styles.markdown} {...props} />,
                              em: ({node, ...props}) => <em className={styles.markdown} {...props} />,
                              code: ({node, ...props}) => <code className={styles.markdown} {...props} />
                            }}
                          >
                            {highlight}
                          </ReactMarkdown>
                        </li>
                      ))}
                    </ul>
                  </div>
                </div>
              </div>
            )}
          </div>
        ))}
      </div>
    </section>
  )
}

export default Education 